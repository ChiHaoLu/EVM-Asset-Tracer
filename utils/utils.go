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
