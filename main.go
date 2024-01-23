package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"github.com/ChiHaoLu/EVM-Asset-Tracer/pkg"
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

	ethPrice, err := pkg.Quote("ETH", "USD")
	if err != nil {
		panic(err)
	}
	usdcPrice, err := pkg.Quote("USDC", "USD")
	if err != nil {
		panic(err)
	}
	usdtPrice, err := pkg.Quote("USDT", "USD")
	if err != nil {
		panic(err)
	}
	daiPrice, err := pkg.Quote("DAI", "USD")
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(data.Chain); i++ {
		chainValue := new(big.Float)
		chainName, nativeTokenName, err := pkg.ExtractNetwork(data.Chain[i])
		fmt.Println("Chain: ", chainName)
		if err != nil {
			log.Fatal(err)
		}
		if chainName == "starknet" {
			url := data.Chain[i] + os.Getenv("API_KEY")
			for j := 0; j < len(data.Address); j++ {
				addressValue := new(big.Float)
				account := data.Address[j]
				if len(account) < 60 {
					continue
				}
				fmt.Println("    Account: ", account)
				nativeBal, err := pkg.GetSNTokenBalanceAndValue(url, "0x049d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7", account, "balanceOf")
				if err == nil {
					nativeValue := new(big.Float).Mul(big.NewFloat(ethPrice), nativeBal)
					addressValue = new(big.Float).Add(addressValue, nativeValue)
					fmt.Printf("	- %s Balance: %f    -> Value: %f\n", nativeTokenName, nativeBal, nativeValue)
				}

				usdcBal, err := pkg.GetSNTokenBalanceAndValue(url, "0x053c91253bc9682c04929ca02ed00b3e423f6710d2ee7e0d5ebb06f3ecf368a8", account, "balanceOf")
				if err == nil {
					usdcValue := new(big.Float).Mul(big.NewFloat(usdcPrice), usdcBal)
					addressValue = new(big.Float).Add(addressValue, usdcValue)
					fmt.Printf("	- USDC Balance: %f    -> Value: %f\n", usdcBal, usdcValue)
				}

				// DAI "0x00da114221cb83fa859dbdb4c44beeaa0bb37c7537ad5ae66fe5e0efd20e6eb3"
				chainValue = new(big.Float).Add(addressValue, chainValue)
			}
			fmt.Printf("Chain Value:%f\n\n", chainValue)
			continue
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
			if len(data.Address[j]) > 60 {
				continue
			}
			addressValue := new(big.Float)
			account := common.HexToAddress(data.Address[j])
			fmt.Println("    Account: ", account)

			nativeBal, err := pkg.GetNativeTokenBalance(client, account)
			if err == nil {
				var nativePrice float64
				if nativeTokenName == "ETH" {
					nativePrice = ethPrice
				} else{
					nativePrice, err = pkg.Quote(nativeTokenName, "USD")
				}
				if err != nil {
					panic(err)
				}
				nativeValue := new(big.Float).Mul(big.NewFloat(nativePrice), nativeBal)
				addressValue = new(big.Float).Add(addressValue, nativeValue)
				fmt.Printf("	- %s Balance: %f    -> Value: %f\n", nativeTokenName, nativeBal, nativeValue)
			}

			usdtBal, err := pkg.GetTokenBalance(client, account, "0xdAC17F958D2ee523a2206206994597C13D831ec7")
			if err == nil {
				usdtValue := new(big.Float).Mul(big.NewFloat(usdtPrice), usdtBal)
				addressValue = new(big.Float).Add(addressValue, usdtValue)
				fmt.Printf("	- USDT Balance: %f    -> Value: %f\n", usdtBal, usdtValue)
			}

			usdcBal, err := pkg.GetTokenBalance(client, account, "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
			if err == nil {
				usdcValue := new(big.Float).Mul(big.NewFloat(usdcPrice), usdcBal)
				addressValue = new(big.Float).Add(addressValue, usdcValue)
				fmt.Printf("	- USDC Balance: %f    -> Value: %f\n", usdcBal, usdcValue)
			}

			daiBal, err := pkg.GetTokenBalance(client, account, "0x6B175474E89094C44Da98b954EedeAC495271d0F")
			if err == nil {
				daiValue := new(big.Float).Mul(big.NewFloat(daiPrice), daiBal)
				addressValue = new(big.Float).Add(addressValue, daiValue)
				fmt.Printf("	- DAI Balance: %f    -> Value: %f\n", daiBal, daiValue)
			}

			fmt.Println("	- Address Value:", addressValue)
			chainValue = new(big.Float).Add(addressValue, chainValue)

			pkg.ProduceMDTable(writer, chainName, account, nativeBal, usdtBal, usdcBal, daiBal, addressValue)
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
