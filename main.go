package main

import (
"fmt"
"github.com/ethereum/go-ethereum/ethclient"
"log"
)

const INFURA_URL="https://gorli.infura.io/v3/74ce7b1c7a104effb6ab0b86ff09eaf0"

func main() {
	client, err := ethclient.Dial(INFURA_URL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections

}
