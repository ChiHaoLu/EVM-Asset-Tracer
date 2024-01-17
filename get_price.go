package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/ChiHaoLu/EVM-Asset-Tracer/utils"
)

type CoinGeckoQuote struct {
	ID    string  `json:"id"`
	Price float64 `json:"current_price"`
}

func Quote(apiUrl string, targetToken string) (float64, error) {
	u, err := url.Parse(apiUrl)
	if err != nil {
		panic(err)
	}
	u, _ = url.Parse(u.String() + "/coins/markets")

	base, err := convertAssetSymbol(targetToken)
	if err != nil {
		panic(err)
	}
	quote, err := convertAssetSymbol("USD")
	if err != nil {
		panic(err)
	}

	ps := url.Values{}
	ps.Add("vs_currency", quote)
	ps.Add("ids", base)

	u.RawQuery = ps.Encode()

	var quotes []CoinGeckoQuote
	res, err := http.Get(u.String())
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(res.Body).Decode(&quotes); err != nil {
		panic(err)
	}

	results := utils.Filter(quotes, func(q CoinGeckoQuote) bool {
		return q.ID == base
	})
	if len(results) == 0 {
		panic(err)
	}

	return results[0].Price, nil
}

func convertAssetSymbol(asset string) (string, error) {
	switch asset {
	case "ETH":
		return "ethereum", nil
	default:
		return "", fmt.Errorf("asset is not supported")
	}
}