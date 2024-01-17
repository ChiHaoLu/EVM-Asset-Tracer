package main

import (
	"context"
	"math"
	"math/big"
	"strconv"

	token "github.com/ChiHaoLu/EVM-Asset-Tracer/contract/usdt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetETHBalance(client *ethclient.Client, account common.Address) (*big.Float, error) {
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
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
