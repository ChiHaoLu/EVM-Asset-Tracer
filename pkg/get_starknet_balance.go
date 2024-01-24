package pkg

import (
	"context"
	"fmt"
	"math/big"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/NethermindEth/starknet.go/utils"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
)

/// @dev the token contract must follow the ERC-20 interface in StarkNet
func GetSNTokenBalanceAndValue(clientUrl string, tokenAddress string, accountAddress string, contractMethod string) (*big.Float, error) {
	c, err := ethrpc.DialContext(context.Background(), clientUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the client, did you specify the url?")
	}
	clientv02 := rpc.NewProvider(c)
	tokenAddressInFelt, err := utils.HexToFelt(tokenAddress)
	if err != nil {
		return nil, err
	}

	accountAddressInFelt, err := utils.HexToFelt(accountAddress)
	if err != nil {
		return nil, err
	}

	// Make read contract call
	tx := rpc.FunctionCall{
		ContractAddress:    tokenAddressInFelt,
		EntryPointSelector: utils.GetSelectorFromNameFelt(contractMethod),
		Calldata:           []*felt.Felt{accountAddressInFelt},
	}
	callResp, err := clientv02.Call(context.Background(), tx, rpc.BlockID{Tag: "latest"})
	if err != nil {
		return nil, err
	}

	// Get token's decimals
	getDecimalsTx := rpc.FunctionCall{
		ContractAddress:    tokenAddressInFelt,
		EntryPointSelector: utils.GetSelectorFromNameFelt("decimals"),
	}
	getDecimalsResp, err := clientv02.Call(context.Background(), getDecimalsTx, rpc.BlockID{Tag: "latest"})
	if err != nil {
		return nil, err
	}

	return FeltToFloat(callResp[0], utils.FeltToBigInt(getDecimalsResp[0])), nil
}
