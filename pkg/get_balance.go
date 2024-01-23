package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
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

func GetBNBBalance(account common.Address) (*big.Float, error) {
	const rpcEndpoint = "https://bsc-dataseed.binance.org/"

	type RPCRequest struct {
		JsonRPC string      `json:"jsonrpc"`
		Method  string      `json:"method"`
		Params  interface{} `json:"params"`
		ID      int         `json:"id"`
	}

	// Create JSON-RPC request
	requestBody, err := json.Marshal(RPCRequest{
		JsonRPC: "2.0",
		Method:  "eth_getBalance",
		Params:  []interface{}{account, "latest"},
		ID:      1,
	})
	if err != nil {
		return nil, err
	}

	// Make HTTP POST request to BSC node
	resp, err := http.Post(rpcEndpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return nil, err
	}

	// Parse the JSON response
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	// Extract the balance from the response
	result, ok := response["result"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid result")
	}
	decimalBalance := ConvertHexToDecimal(result)
	return decimalBalance, nil
}
