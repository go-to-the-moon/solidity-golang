package src


import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func Client() {
	client, err := ethclient.Dial(INFURA_URL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
}


