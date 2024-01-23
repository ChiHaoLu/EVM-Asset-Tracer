package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"bufio"

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

	report, err := os.Create("BalanceResult.md")
	if err != nil {
		log.Fatal(err)
	}
	defer report.Close()
	writer := bufio.NewWriter(report)
	defer writer.Flush() 
	writer.WriteString("| Chain | Address | ETH Balance | USDT Balance | USDC Balance | DAI Balance | MATIC Balance | BNB Balance | Value |\n")
	writer.WriteString("|-------|---------|-------------|--------------|--------------|-------------|---------------|-------------|---------------|\n")

	allValue := new(big.Float)
	for i := 0; i < len(data.Chain); i++ {
		chainValue := new(big.Float)
		chainName, nativeTokenName, err := utils.ExtractNetwork(data.Chain[i])
		fmt.Println("Chain: ", chainName)
		if err != nil {
			log.Fatal(err)
		}

		var url string
		if chainName == "bsc" {
			url = data.Chain[i]
		} else {
			url = data.Chain[i] + os.Getenv("API_KEY")
		}

		client, err := ethclient.Dial(url)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Close()

		for j := 0; j < len(data.Address); j++ {
			addressValue := new(big.Float)
			account := common.HexToAddress(data.Address[j])
			fmt.Println("    Account: ", account)

			nativeBal, err := utils.GetNativeTokenBalance(client, account)
			if err == nil {
				nativePrice, err := utils.Quote(nativeTokenName, "USD")
				if err != nil {
					panic(err)
				}
				nativeValue := new(big.Float).Mul(big.NewFloat(nativePrice), nativeBal)
				addressValue = new(big.Float).Add(addressValue, nativeValue)
				fmt.Printf("	- %s Balance: %f    -> Value: %f\n", nativeTokenName, nativeBal, nativeValue)
			}

			usdtBal, err := utils.GetTokenBalance(client, account, "0xdAC17F958D2ee523a2206206994597C13D831ec7")
			if err == nil {
				usdtPrice, err := utils.Quote("USDT", "USD")
				if err != nil {
					panic(err)
				}
				usdtValue := new(big.Float).Mul(big.NewFloat(usdtPrice), usdtBal)
				addressValue = new(big.Float).Add(addressValue, usdtValue)
				fmt.Printf("	- USDT Balance: %f    -> Value: %f\n", usdtBal, usdtValue)
			}
			
			usdcBal, err := utils.GetTokenBalance(client, account, "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
			if err == nil {
				usdcPrice, err := utils.Quote("USDC", "USD")
				if err != nil {
					panic(err)
				}
				usdcValue := new(big.Float).Mul(big.NewFloat(usdcPrice), usdcBal)
				addressValue = new(big.Float).Add(addressValue, usdcValue)
				fmt.Printf("	- USDC Balance: %f    -> Value: %f\n", usdcBal, usdcValue)
			}

			daiBal, err := utils.GetTokenBalance(client, account, "0x6B175474E89094C44Da98b954EedeAC495271d0F")
			if err == nil {
				daiPrice, err := utils.Quote("DAI", "USD")
				if err != nil {
					panic(err)
				}
				daiValue := new(big.Float).Mul(big.NewFloat(daiPrice), daiBal)
				addressValue = new(big.Float).Add(addressValue, daiValue)
				fmt.Printf("	- DAI Balance: %f    -> Value: %f\n", daiBal, daiValue)
			}

			fmt.Println("	- Address Value:", addressValue)
			chainValue = new(big.Float).Add(addressValue, chainValue)

			utils.ProduceMDTable(writer, chainName, account, nativeBal, usdtBal, usdcBal, daiBal, addressValue)
		}
		fmt.Printf("Chain Value:%f\n\n", chainValue)
		allValue = new(big.Float).Add(allValue, chainValue)
		fmt.Fprintf(writer, "| %s | %s |  |  |  | |  |  | %f |\n",
			chainName,
			"Total",
			chainValue,
		)
	}
	fmt.Printf("All Value:%f\n", allValue)
	fmt.Fprintf(writer, "| %s | |  |  |  | |  |  | %f |\n",
		"Total",
		allValue,
	)
}
