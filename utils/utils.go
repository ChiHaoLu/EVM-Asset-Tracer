package utils

import (
	"strings"
)

func Filter[T any](elems []T, filter func(elem T) bool) []T {
	var res []T
	for _, elem := range elems {
		if filter(elem) {
			res = append(res, elem)
		}
	}
	return res
}

func ExtractNetwork(url string) (string, string, string, error) {
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
	var wethAddress string
	if networkName == "polygon" {
		nativeTokenName = "MATIC"
		wethAddress = "0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619"
	} else if networkName == "optimism" {
		nativeTokenName = "OP"
		wethAddress = "0x4200000000000000000000000000000000000006"
	} else if networkName == "arbitrum" {
		nativeTokenName = "ARB"
		wethAddress = "0x82aF49447D8a07e3bd95BD0d56f35241523fBab1"
	} else if networkName == "bnbsmartchain" {
		nativeTokenName = "BNB"
		wethAddress = "0x4DB5a66E937A9F4473fA95b1cAF1d1E1D62E29EA"
	} else {
		nativeTokenName = "ETH"
		wethAddress = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
	}
	return networkName, nativeTokenName, wethAddress, nil
}
