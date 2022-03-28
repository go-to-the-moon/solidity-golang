package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
	"log"
)

var INFURA_URL = "https://gorli.infura.io/v3/74ce7b1c7a104effb6ab0b86ff09eaf0"

var config *viper.Viper
var AllConfig TotalConfig

type Infura struct {
	Url string
}

type TotalConfig struct {
	Infura Infura
}


func main() {

	config = viper.New()
	// 设置 public 配置文件名
	config.SetConfigFile("./config/config.yaml")
	// 读取该配置文件
	config.ReadInConfig()

	// 解析 secret config
	config.SetConfigFile("./config/config.yaml")
	config.MergeInConfig()

	var  url =  config.GetString("infura")

	fmt.Println("current infura_url:",url)



	client, err := ethclient.Dial(INFURA_URL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections

}
