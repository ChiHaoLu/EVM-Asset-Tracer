package constant

type AddressList struct {
	mainnet       string
	polygon       string
	optimism      string
	arbitrum      string
	bnbsmartchain string
	starknet      string
}

func GetWETHAddressFromArray() AddressList {
	return AddressList{
		mainnet:       "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
		polygon:       "0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619",
		optimism:      "0x4200000000000000000000000000000000000006",
		arbitrum:      "0x82aF49447D8a07e3bd95BD0d56f35241523fBab1",
		bnbsmartchain: "0x4DB5a66E937A9F4473fA95b1cAF1d1E1D62E29EA",
		starknet:      "",
	}
}

func GetARBAddressFromArray() AddressList {
	return AddressList{
		mainnet:       "0xB50721BCf8d664c30412Cfbc6cf7a15145234ad1",
		polygon:       "",
		optimism:      "",
		arbitrum:      "0x912CE59144191C1204E64559FE8253a0e49E6548",
		bnbsmartchain: "",
		starknet:      "",
	}
}

func GetOPAddressFromArray() AddressList {
	return AddressList{
		mainnet:       "",
		polygon:       "",
		optimism:      "0x4200000000000000000000000000000000000042",
		arbitrum:      "",
		bnbsmartchain: "",
		starknet:      "",
	}
}

func GetSTRKAddressFromArray() AddressList {
	return AddressList{
		mainnet:       "0xCa14007Eff0dB1f8135f4C25B34De49AB0d42766",
		polygon:       "",
		optimism:      "",
		arbitrum:      "",
		bnbsmartchain: "",
		starknet:      "0x04718f5a0fc34cc1af16a1cdee98ffb20c31f5cd61d6ab07201858f4287c938d",
	}
}
