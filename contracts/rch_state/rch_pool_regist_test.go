package rch_state

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestName(t *testing.T) {

	url := "http://47.57.138.200:8488"

	http, err := ethclient.Dial(url)

	if err != nil {
		panic(err)
	}

	interfaceAddr := common.HexToAddress(Rch_contract_interface_addr)
	instance, err := NewRchMgrInterface(interfaceAddr, http)
	if err != nil {
		panic(err)
	}

	rchRegPoolMgrAddr, er1 := instance.GetContractAddr(nil, "RchRegPoolMgr")
	if er1 != nil {
		panic(er1)
	}

	fmt.Println(rchRegPoolMgrAddr.Hex())

}
