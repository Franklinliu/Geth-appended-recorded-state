package vm

import (
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

type opFunc func(pc *uint64, evm *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error)

func Hacker_record(op OpCode, fun opFunc, pc *uint64, evm *EVM, contract *Contract, memory *Memory, stack *Stack) ([]byte, error) {
	var code_desc string = strconv.FormatUint(*pc, 10) + opCodeToString[op]
	if GetGlobalWatchDog().TurnOn() == true {
		GetGlobalWatchDog().Write2Trace(code_desc)
		GetGlobalWatchDog().GetEnv().StateDB.ForEachStorage(*(GetGlobalWatchDog().GetTx().To()), func(key, value common.Hash) bool {
			GetGlobalWatchDog().Write2Storage(key, value)
			return true
		})
	}
	if GetGlobalTracerWatchDog().TurnOn() == true {
		GetGlobalTracerWatchDog().Write2Trace(code_desc)
		GetGlobalTracerWatchDog().GetEnv().StateDB.ForEachStorage(*(GetGlobalTracerWatchDog().GetTx().To()), func(key, value common.Hash) bool {
			GetGlobalTracerWatchDog().Write2Storage(key, value)
			return true
		})
	}
	return fun(pc, evm, contract, memory, stack)
}
