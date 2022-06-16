package etxhash

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"math"
	"net"
	"strings"
	//"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math/big"
	"time"
)


const (
	MINER_RATE2        = 0.51
	NODE_RATE2         = 0.39
	COMMUNITY_RATE2    = 0.10
	ETX_LIMIT_BLOCK_V1 = 500
	ETX_LIMIT_BLOCK_V2 = 666666
	//MINER_RATE3        = 0.60
	//NODE_RATE3        = 0.30
	//COMMUNITY_RATE3    = 0.10
	//ETX_LIMIT_BLOCK_V2 = 20000
	ETX_BRANCH_BLOCK_V1 = 56500
	BLOCK_REWARD = 833
	ETHV_REDUCE_BOUNDING = 1110858  // 半年减半机制
	Version            = "RCH v3.0"
)

var (
	big8  = big.NewInt(8)
	big32 = big.NewInt(32)
)
// Ethash proof-of-work protocol constants.
var (
	FrontierBlockReward       = big.NewInt((0e+18)) // Block reward in wei for successfully mining a block changed .
	ByzantiumBlockReward      = big.NewInt(0e+18) // Block reward in wei for successfully mining a block upward from Byzantium
	ConstantinopleBlockReward = big.NewInt(0e+18) // Block reward in wei for successfully mining a block upward from Constantinople
	maxUncles                 = 2                 // Maximum number of uncles allowed in a single block
	allowedFutureBlockTime    = 15 * time.Second  // Max time from current time allowed for blocks, before they're considered future blocks

	// added .
	eth1 = big.NewInt(1e18)
	StartMineTime = time.Date(2020,10, 10, 8, 8, 8, 8, time.UTC)
	//减半机制
	//Year1_BlockReward = []*big.Int {eth1.Mul(eth1, big.NewInt(10)), big.NewInt((10)), big.NewInt((5e+18)),
	//	big.NewInt((5e+18)), big.NewInt((2.5e+18)), big.NewInt((2.5e+18)),  }
	TotalAccounts = map[string]*big.Int{}
	CurrentTotalAccounts = 0
	LowBalance = false
	coinBase common.Address
	coinBase2 common.Address
	//0x6a05b4fc80659715079204df739f5f61288a13e5
	nodeAddress1      = common.HexToAddress("0xb5bEc342E0a170b7f779d6504bb47f689D93c96d")  // 39%
	communityAddress1      = common.HexToAddress("0x75e9bD9b37F4aA381e1e679Cd03E319b56353d6e")  //10 %

)


var rpcLocalHost = "http://127.0.0.1:8488"
var rpcETX1Host = "http://127.0.0.1:8488"

var TotalNodes = 0

var blackList = []string {
	"0xEd8c90ef765BD7fCbDd3559E2B37549A91059D00", //test1
	"0x7a737683ff97790c8b717339a8b97a897b3ab5c4", //fubt
}

var whiteList_fix = []string {
	//"0x47c25d26431ea4e7fa15f6a1dcbd6331d44e6617", // rch
	"0xd043A94Db9F5623790644B26833DeC93A8422269",
	"0xe1aFC49aad79e778735772ea12b7Da1DDC1A6733", // rch new pool
	"0xb5225f2a4f4B8cFc16086B23Df1Ca152106A85D9", // test.
}

var whiteList = []string {

}

type CoinReward struct {
	Amount decimal.Decimal
	StartTime   int64
	UpdateTime  int64
}

type PoolStat struct {
	Coins24 map[string]*CoinReward
	StartTime   int64
	UpdateTime  int64
}

var whiteListEnableV6 = false
var gpoolInfo PoolStat


func AccumulateRewardsV5(config *params.ChainConfig, state *state.StateDB, header *types.Header, uncles []*types.Header) {
	// Select the correct block reward based on chain progression
	//eth_1 := big.NewInt(1e18)
	//eth_stake_limit := eth_1.Mul(eth_1, big.NewInt(10000)) //9999

	blockReward := FrontierBlockReward
	if config.IsByzantium(header.Number) {
		blockReward = ByzantiumBlockReward
	}

	if config.IsConstantinople(header.Number) {
		blockReward = ConstantinopleBlockReward
	}

	//coinBaseBalance:= state.GetBalance(coinBase2)
	//if header.Number.Cmp(big.NewInt(ETX_LIMIT_BLOCK_V1))> 0  &&  coinBase2.Hex()!= "0x0000000000000000000000000000000000000000" &&
	//	coinBaseBalance.Cmp(eth_stake_limit) <0  {
	//	log.Println("stake not enough , unable to mine v4.", coinBase2.Hex(), coinBaseBalance.String(), eth_stake_limit.String())
	//	LowBalance = true
	//}else {
	//	LowBalance = false
	//}

	baseBalance:= state.GetBalance(header.Coinbase)
	//log.Println("抵押挖矿", header.Coinbase.Hex(), baseBalance.String(), eth_stake_value.String())

	headerTime:= time.Unix( int64(header.Time),0 )
	if baseBalance.Cmp(big.NewInt(0)) > 0 {
		TotalAccounts[header.Coinbase.String()] = baseBalance
		CurrentTotalAccounts = len(TotalAccounts)
	}

	reduceMod:= header.Number.Int64()/ETHV_REDUCE_BOUNDING
	reward1:= decimal.New(BLOCK_REWARD, 18)
	blockReward = reward1.Mul(decimal.NewFromFloat((math.Pow(0.5, float64(reduceMod))))).BigInt()
	currReward:= decimal.NewFromBigInt(blockReward, -18)
	log.Println("mine start time：", StartMineTime.Local(), Version, headerTime, LowBalance, currReward.String(), "nodes:", CurrentTotalAccounts)

	var blockReward4 = big.NewInt(0)

	temp_blockReward:= *blockReward
	fblockReward:= float64(temp_blockReward.Uint64())

	if header.Number.Cmp(big.NewInt(ETX_LIMIT_BLOCK_V1)) <= 0 {
		fblockReward = 0
	}

	if header.Number.Cmp(big.NewInt(ETX_LIMIT_BLOCK_V2)) <= 0 {
		fblockReward4 := fblockReward*1
		blockReward4 = big.NewInt(int64(fblockReward4))
		log.Println("Miner reward：", 1, blockReward4)

	}else {
		fblockReward3 := fblockReward*NODE_RATE2
		iteamReward3 := big.NewInt(int64(fblockReward3))
		state.AddBalance(nodeAddress1, iteamReward3)
		//state.AddBalance(nodeAddress1, iteamReward3)
		log.Println("节点奖励：", NODE_RATE2, iteamReward3.String())

		fblockReward2 := fblockReward*COMMUNITY_RATE2
		iteamReward2 := big.NewInt(int64(fblockReward2))
		state.AddBalance(communityAddress1, iteamReward2)
		log.Println("基金奖励：", COMMUNITY_RATE2, iteamReward2.String())

		fblockReward4 := fblockReward*MINER_RATE2
		blockReward4 = big.NewInt(int64(fblockReward4))
		log.Println("Miner reward2：", MINER_RATE2, blockReward4)
	}

	// Accumulate the rewards for the miner and any included uncles
	reward := new(big.Int).Set(blockReward4)
	r := new(big.Int)
	for _, uncle := range uncles {
		r.Add(uncle.Number, big8)
		r.Sub(r, header.Number)
		r.Mul(r, blockReward4)
		r.Div(r, big8)
		//uncleBaseBalance:= state.GetBalance(uncle.Coinbase)
		//if header.Number.Cmp(big.NewInt(ETX_LIMIT_BLOCK_V1))> 0 &&  uncle.Coinbase.Hex()!= "0x0000000000000000000000000000000000000000" &&
		//	uncleBaseBalance.Cmp(eth_stake_limit) <0  {
		//	log.Println("stake not enough for uncle.", uncle.Coinbase.Hex(), uncleBaseBalance.String(), eth_stake_limit.String())
		//	//blockReward = big.NewInt(0)
		//}
		reward.Add(reward, r)
	}

	state.AddBalance(header.Coinbase, reward)

}


func AccumulateRewardsV6(config *params.ChainConfig, state *state.StateDB, header *types.Header, uncles []*types.Header) {
	// Select the correct block reward based on chain progression
	//eth_1 := big.NewInt(1e18)
	//eth_stake_limit := eth_1.Mul(eth_1, big.NewInt(10000)) //9999

	blockReward := FrontierBlockReward
	if config.IsByzantium(header.Number) {
		blockReward = ByzantiumBlockReward
	}

	if config.IsConstantinople(header.Number) {
		blockReward = ConstantinopleBlockReward
	}

	//coinBaseBalance:= state.GetBalance(coinBase2)
	//if header.Number.Cmp(big.NewInt(ETX_LIMIT_BLOCK_V1))> 0  &&  coinBase2.Hex()!= "0x0000000000000000000000000000000000000000" &&
	//	coinBaseBalance.Cmp(eth_stake_limit) <0  {
	//	log.Println("stake not enough , unable to mine v4.", coinBase2.Hex(), coinBaseBalance.String(), eth_stake_limit.String())
	//	LowBalance = true
	//}else {
	//	LowBalance = false
	//}

	baseBalance:= state.GetBalance(header.Coinbase)
	//log.Println("抵押挖矿", header.Coinbase.Hex(), baseBalance.String(), eth_stake_value.String())

	headerTime:= time.Unix( int64(header.Time),0 )
	if baseBalance.Cmp(big.NewInt(0)) > 0 {
		TotalAccounts[header.Coinbase.String()] = baseBalance
		CurrentTotalAccounts = len(TotalAccounts)
	}

	reduceMod:= header.Number.Int64()/ETHV_REDUCE_BOUNDING
	reward1:= decimal.New(BLOCK_REWARD, 18)
	blockReward1 := reward1.Mul(decimal.NewFromFloat((math.Pow(0.5, float64(reduceMod)))))
	blockReward.SetString(blockReward1.String(), 10)
	currReward:= decimal.NewFromBigInt(blockReward, -18)
	log.Println("mine start time：", StartMineTime.Local(), Version, headerTime, LowBalance, currReward.String(), "nodes:", CurrentTotalAccounts)

	var blockReward4 = big.NewInt(0)

	//temp_blockReward:= *blockReward
	//fblockReward:= float64(temp_blockReward.Uint64())

	if header.Number.Cmp(big.NewInt(ETX_LIMIT_BLOCK_V1)) <= 0 {
		blockReward1 = decimal.NewFromFloat(0)
	}

	if header.Number.Cmp(big.NewInt(ETX_LIMIT_BLOCK_V2)) <= 0 {
		fblockReward4:= blockReward1.Mul(decimal.NewFromFloat(1.0))
		//fblockReward4 := fblockReward*1
		blockReward4 = big.NewInt(0)//big.NewInt(int64(fblockReward4))
		blockReward4.SetString(fblockReward4.String(), 10 )
		reward4:= decimal.NewFromBigInt(blockReward4, -18)
		log.Println("Miner reward：", 1, reward4.String())

	}else {
		fblockReward3:= blockReward1.Mul(decimal.NewFromFloat(NODE_RATE2))
		//fblockReward3 := fblockReward*NODE_RATE2
		iteamReward3 := big.NewInt(0)
		iteamReward3.SetString(fblockReward3.String(), 10)

		state.AddBalance(nodeAddress1, iteamReward3)
		//state.AddBalance(nodeAddress1, iteamReward3)
		log.Println("NODE：", NODE_RATE2, fblockReward3.Mul(decimal.New(1, -18)).String())

		fblockReward2 := blockReward1.Mul(decimal.NewFromFloat(COMMUNITY_RATE2))
		iteamReward2 := big.NewInt(0)
		iteamReward2.SetString(fblockReward2.String(), 10)
		state.AddBalance(communityAddress1, iteamReward2)
		log.Println("COMMUNITY：", COMMUNITY_RATE2, fblockReward2.Mul(decimal.New(1, -18)).String())

		fblockReward4 := blockReward1.Mul(decimal.NewFromFloat(MINER_RATE2)) //fblockReward*MINER_RATE2
		blockReward4 = big.NewInt(0)
		blockReward4.SetString(fblockReward4.String(), 10)
		//log.Println("Miner reward2：", MINER_RATE2, blockReward4)
		//reward4:= decimal.NewFromBigInt(blockReward4, -18)
		log.Println("Miner reward2：", 1, fblockReward4.Mul(decimal.New(1, -18)).String())
	}

	// Accumulate the rewards for the miner and any included uncles
	reward := new(big.Int).Set(blockReward4)
	r := new(big.Int)
	for _, uncle := range uncles {
		r.Add(uncle.Number, big8)
		r.Sub(r, header.Number)
		r.Mul(r, blockReward4)
		r.Div(r, big8)
		//uncleBaseBalance:= state.GetBalance(uncle.Coinbase)
		//if header.Number.Cmp(big.NewInt(ETX_LIMIT_BLOCK_V1))> 0 &&  uncle.Coinbase.Hex()!= "0x0000000000000000000000000000000000000000" &&
		//	uncleBaseBalance.Cmp(eth_stake_limit) <0  {
		//	log.Println("stake not enough for uncle.", uncle.Coinbase.Hex(), uncleBaseBalance.String(), eth_stake_limit.String())
		//	//blockReward = big.NewInt(0)
		//}
		//reward.Add(reward, r)
	}

	reward5:= decimal.NewFromBigInt(reward, 0)
	log.Println("Miner reward3：", 1, header.Coinbase.String(), reward5.String())
	state.AddBalance(header.Coinbase, reward)

}


func SetCoinBase(addr common.Address) {
	coinBase2 = addr
}


func IsMakeWork()bool {
	//log.Println("IsMakeWork:", LowBalance)
	return true
	return !LowBalance
}

func GetLocalMacAddr() string {
	interfaces,err := net.Interfaces()
	if err != nil {
		log.Println("Error : " + err.Error())
		return ""
	}

	for _,inter := range interfaces {
		mac := inter.HardwareAddr //获取本机MAC地址
		log.Println("MAC = ",mac, inter)
		if inter.Name == "en0" {
			return mac.String()
		}
	}

	return ""
}


func Is_black_list(ac string )(bool ){
	for _, i:= range blackList {
		if strings.ToUpper(i) == strings.ToUpper(ac ) {
			return true
		}
	}
	return false ;
}

func SetWhiteList(addrs []string) {
	whiteList = whiteList[:0]
	whiteList = append(whiteList, whiteList_fix...)
	for _, it:= range addrs {
		whiteList = append(whiteList, it)

	}
}

func Is_whitelist_v6()(bool ){

	return true ;

	for _, i:= range whiteList {
		if strings.ToUpper(i) == strings.ToUpper(coinBase2.Hex() ) {
			return true
		}
	}

	//type Callback func()(ret string)
	//var pointer *Callback
	//var getPoolList Callback = rch_state.GetPoolList
	//straddress := &(getPoolList)
	//pointer = *(**Callback)(unsafe.Pointer(&straddress))
	//
	//fmt.Println( (Callback)(*pointer))

	return false ;

}