package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"github.com/ChiHaoLu/EVM-Asset-Tracer/utils"
)

type Metadata struct {
	Address []string `json:"address"`
	Chain   []string `json:"chain"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	var data Metadata
	file, err := os.ReadFile("metadata.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(data.Chain); i++ {
		chainName, err := utils.ExtractNetworkName(data.Chain[i])
		fmt.Println("Chain: ", chainName)
		if err != nil {
			log.Fatal(err)
		}

		client, err := ethclient.Dial(data.Chain[i] + os.Getenv("API_KEY"))
		if err != nil {
			log.Fatal(err)
		}
		defer client.Close()

		for j := 0; j < len(data.Address); j++ {
			account := common.HexToAddress(data.Address[j])
			fmt.Println("    Account: ", account)

			ethBal, err := GetETHBalance(client, account)
			if err == nil {
				fmt.Println("	- ETH Balance:", ethBal)
			}

			usdtBal, err := GetTokenBalance(client, account, "0xdAC17F958D2ee523a2206206994597C13D831ec7")
			if err == nil {
				fmt.Println("	- USDT Balance:", usdtBal)
			}

			daiBal, err := GetTokenBalance(client, account, "0x6B175474E89094C44Da98b954EedeAC495271d0F")
			if err == nil {
				fmt.Println("	- DAI Balance:", daiBal)
			}

			usdcBal, err := GetTokenBalance(client, account, "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
			if err == nil {
				fmt.Println("	- USDC Balance:", usdcBal)
			}
		}
		fmt.Printf("\n")
	}
}
