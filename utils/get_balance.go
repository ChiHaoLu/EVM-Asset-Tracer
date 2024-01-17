package utils

import (
	"math"
	"math/big"
	"strconv"

	token "github.com/ChiHaoLu/EVM-Asset-Tracer/contract/usdt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetNativeTokenBalance(client *ethclient.Client, account common.Address) (*big.Float, error) {
	rpcclient := client.Client()

	var result string
	if err := rpcclient.Call(&result, "eth_getBalance", account, "latest"); err != nil {
		return nil, err
	}
	fbalance := new(big.Float)
	fbalance.SetString(result)
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	return ethValue, nil
}

func GetTokenBalance(client *ethclient.Client, account common.Address, tokenAddress string) (*big.Float, error) {
	address := common.HexToAddress(tokenAddress)
	instance, err := token.NewUsdt(address, client)
	if err != nil {
		return nil, err
	}

	bal, err := instance.BalanceOf(&bind.CallOpts{}, account)
	if err != nil {
		return nil, err
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	decimalInStr := decimals.String()
	decimalInInt, _ := strconv.Atoi(decimalInStr)

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(decimalInInt)))

	return value, nil
}
