package utils

import (
	"fmt"
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

func ExtractNetworkName(url string) (string, error) {
	p1 := strings.Split(url, "//")
	p2 := strings.Split(p1[1], "-")
	if len(p2) < 2 {
		p2 = strings.Split(p1[1], ".")
		if len(p2) < 2 {
			return "", fmt.Errorf("chain url is invalid")
		}
	}
	networkName := p2[0]
	return networkName, nil
}