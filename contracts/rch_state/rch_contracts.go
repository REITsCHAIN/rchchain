package rch_state

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

//const  rpcTestRchHost = "http://127.0.0.1:8488"

const  Rch_contract_interface_addr = "0xcFd6cD2C5F7C519a9C8f4A4Eb70dfa6D48967A64"

type RegPoolStatus  struct {
	UserAddr    common.Address
	UserBalane  *big.Int
	UpdateTime  *big.Int
	Unstakeing  bool
	UnstakeGate *big.Int
}

var rchRegPoolStatus map[common.Address]bool = make(map[common.Address]bool)

var eth_backend bind.ContractBackend

func SetRchBackend(backend bind.ContractBackend){
	eth_backend = backend
}

func CheckPoolValidation(addr string) bool {
	add1:= common.HexToAddress(addr)
	return rchRegPoolStatus[add1]

}

func GetRegPoolList()(ret []string, er error){

	interfaceAddr := common.HexToAddress(Rch_contract_interface_addr)
	instance, err := NewRchMgrInterface(interfaceAddr, eth_backend)
	if err != nil {
		return nil, err
	}

	rchRegPoolMgrAddr, er1 := instance.GetContractAddr(nil, "RchRegPoolMgr")
	if er1!= nil {
		return  nil, er1
	}

	RchRegPoolMgr, er2:= NewRchRegPoolMgr(rchRegPoolMgrAddr, eth_backend)
	if er2!= nil {
		return nil, er2
	}

	pool_count, er3 := RchRegPoolMgr.PoolListCount(nil)
	if er3!= nil {
		return nil, er3
	}

	stakegate, er2 := RchRegPoolMgr.GetStakeGateValue(nil)
	//stakegate:= decimal.NewFromFloat(10000*1e18)
	if er2!= nil {
		return nil, er2
	}

	for i:=0; i< (int)(pool_count.Int64()) ; i++ {
		addr, err := RchRegPoolMgr.PoolAddrList(nil , big.NewInt(int64(i)))
		if err != nil {
			break;
		}

		stakeStatus, er3:= RchRegPoolMgr.PoolStakeList(nil, addr )
		if er3!= nil {
			break
		}

		//log.Println("GetPoolList:", addr.Hex(), stakeStatus.UserBalane.String(), stakegate.String())
		if stakeStatus.UserBalane.Cmp(stakegate) >= 0 {
			rchRegPoolStatus[addr] = true
			ret = append(ret, addr.Hex())
		}else {
			rchRegPoolStatus[addr] = false
		}

		//log.Println("GetPoolList:", addr.String(), rchRegPoolStatus[addr])
	}

	return ret, nil ;
}

