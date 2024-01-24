# EVM Compatible Chains and StarkNet Asset Tracer

## Environment
- Go: 1.21.0

## Supports

### Chain

- Mainnet
- Polygon
- Optimism
- Arbitrum
- BSC
- StarkNet

### Token

- Native Token
  - ETH
  - MATIC
  - BNB
- ERC-20
  - ~WETH~
  - ~OP~
  - ~ARB~
  - USDT
  - USDC
  - DAI

## Run

Make sure you have build `abigen` tool from `go-ethereum`.
```
go get -u github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make devtools
./abigen --help
>
NAME:
   abigen - Ethereum ABI wrapper code generator

USAGE:
   abigen [global options] command [command options] [arguments...]

VERSION:
   1.13.3-unstable-052355f5-20231004

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   
    --abi value                                                           
          Path to the Ethereum contract ABI json to bind, - for STDIN
   
    --alias value                                                         
          Comma separated aliases for function and event renaming, e.g. original1=alias1,
          original2=alias2
   
    --bin value                                                           
          Path to the Ethereum contract bytecode (generate deploy method)
   
    --combined-json value                                                 
          Path to the combined-json file generated by compiler, - for STDIN
   
    --exc value                                                           
          Comma separated types to exclude from binding
   
    --lang value                        (default: "go")                   
          Destination language for the bindings (go)
   
    --out value                                                           
          Output file for the generated binding (default = stdout)
   
    --pkg value                                                           
          Package name to generate the binding into
   
    --type value                                                          
          Struct name for the binding (default = package name)

   MISC

   
    --help, -h                          (default: false)                  
          show help
   
    --version, -v                       (default: false)                  
          print the version


COPYRIGHT:
   Copyright 2013-2023 The go-ethereum Authors
```

Produce all token's go contract with below command
```
$ SOLC_VERSION=0.4.17 solc --version
SOLC_VERSION=0.4.17 solc --abi TetherToken.sol --out abi
SOLC_VERSION=0.4.17 solc --bin TetherToken.sol --out bin
abigen --bin=bin/TetherToken.bin --abi=build/TetherToken.abi --pkg=usdt --out=usdt.go --alias _totalSupply=TotalSupply1
```

Run below commands to get the accounts' total balance
```
$ go get <gihub.com/repo_org/repo_name> // If you want to add new package
$ go mod download
$ go run *.go
```

## Appendic
1. `eth_getBalance` only returns the balance of the native chain currency (ex: ETH for Ethereum or Matic for Polygon) and does not include any ERC20 token balances for the given address. [ref.](https://docs.alchemy.com/reference/eth-getbalance-polygon)