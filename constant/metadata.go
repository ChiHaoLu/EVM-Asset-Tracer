package constant

const (
    ADDRESS_LEN = 4
    CHAIN_LEN = 6
)

func GetAddressFromArray(n int) string {
	return [...]string{
		"0x2b83c71a59b926137d3e1f37ef20394d0495d72d",
		"0x189C92f28047c979cA2D17C13e3A12963EB1b8B4",
		"0xd837854e837F7011300fB422748fb40Cc4D3E8bc",
		"0x03042603baBd486b6a52CB56BA2b4533C294D84b14FF84879BBd4BdEda16B417",
	}[n]
}

func GetCahinUrlPrefixFromArray(n int) string {
	return [...]string{
		"https://mainnet.infura.io/v3/",
		"https://polygon-mainnet.infura.io/v3/",
		"https://optimism-mainnet.infura.io/v3/",
		"https://arbitrum-mainnet.infura.io/v3/",
		"https://bsc-dataseed.binance.org/",
		"https://starknet-mainnet.infura.io/v3/",
	}[n]
}
