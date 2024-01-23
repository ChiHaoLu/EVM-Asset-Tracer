package pkg

import (
	"bufio"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func ConvertHexToDecimal(hexValue string) *big.Float {
	intValue, success := new(big.Int).SetString(hexValue[2:], 16)
	if !success {
		return nil
	}
	floatValue := new(big.Float).SetInt(intValue)
	return floatValue
}

func Filter[T any](elems []T, filter func(elem T) bool) []T {
	var res []T
	for _, elem := range elems {
		if filter(elem) {
			res = append(res, elem)
		}
	}
	return res
}

func ExtractNetwork(url string) (string, string, error) {
	p1 := strings.Split(url, "//")
	p2 := strings.Split(p1[1], "-")
	if len(p2) < 2 {
		p2 = strings.Split(p1[1], ".")
		if len(p2) < 2 {
			panic("chain url is invalid")
		}
	}
	networkName := p2[0]

	var nativeTokenName string
	if networkName == "polygon" {
		nativeTokenName = "MATIC"
	} else if networkName == "bsc" {
		nativeTokenName = "BNB"
	} else {
		nativeTokenName = "ETH"
	}
	return networkName, nativeTokenName, nil
}

func ProduceMDTable(writer *bufio.Writer,
	chainName string,
	account common.Address,
	nativeBal *big.Float,
	usdtBal *big.Float,
	usdcBal *big.Float,
	daiBal *big.Float,
	addressValue *big.Float) {

	zeroValue := new(big.Float)
	if nativeBal == nil {
		nativeBal = zeroValue
	}
	if usdtBal == nil {
		usdtBal = zeroValue
	}
	if usdcBal == nil {
		usdcBal = zeroValue
	}
	if daiBal == nil {
		daiBal = zeroValue
	}

	if chainName == "polygon" {
		fmt.Fprintf(writer, "| %s | %s | %f | %f | %f | %f | %f | %f | %f |\n",
			chainName,
			account,
			zeroValue,
			usdtBal,
			usdcBal,
			daiBal,
			nativeBal,
			zeroValue,
			addressValue,
		)
	} else if chainName == "bsc" {
		fmt.Fprintf(writer, "| %s | %s | %f | %f | %f | %f | %f | %f | %f |\n",
			chainName,
			account,
			zeroValue,
			usdtBal,
			usdcBal,
			daiBal,
			zeroValue,
			nativeBal,
			addressValue,
		)
	} else {
		fmt.Fprintf(writer, "| %s | %s | %f | %f | %f | %f | %f | %f | %f |\n",
			chainName,
			account,
			nativeBal,
			usdtBal,
			usdcBal,
			daiBal,
			zeroValue,
			zeroValue,
			addressValue,
		)
	}
}
