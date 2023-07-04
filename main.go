package main

import (
	"fmt"

	"github.com/bettybao1209/lingjing-sdk/invoke"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetEthClient(chainUrl string) *ethclient.Client {
	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		panic(err)
	}
	return client
}

func main() {
	rpc := "https://test-node.lingjing-eco.com.cn/"
	contract := "0x49d2a58474e8ceb5ec91d1784f5abc30ae5adc3b"
	client := GetEthClient(rpc)

	if name, err := invoke.GetName(client, contract); err == nil {
		fmt.Println(name)
	}
}
