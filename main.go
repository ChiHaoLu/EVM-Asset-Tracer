package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

type Metadata struct {
	Address []string `json:"address"`
	Chain   []string `json:"chain"`
	Token   []string `json:"token"`
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
		client, err := ethclient.Dial(data.Chain[i] + os.Getenv("API_KEY"))
		if err != nil {
			log.Fatal(err)
		}
		defer client.Close()

		for j := 0; j < len(data.Address); j++ {
			account := common.HexToAddress(data.Address[j])
			ethBal, err := GetETHBalance(client, account)
			if err == nil {
				fmt.Println("ETH Balance:", ethBal)
			}

			for k := 0; k < len(data.Token); k++ {

			}
		}
	}
}
