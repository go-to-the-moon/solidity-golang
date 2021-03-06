package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const INFURA_URL = "https://gorli.infura.io/v3/74ce7b1c7a104effb6ab0b86ff09eaf0"

func main() {
	client, err := ethclient.Dial(INFURA_URL)
	if err != nil {
		log.Fatal(err)
	}

	// 账户余额
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 25893180161173005034

	// 区块高度
	blockNumber := big.NewInt(5532993)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt) // 25729324269165216042

	// ETH余额
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041

	// pending余额
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance) // 25729324269165216042
}