/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

var (
	//// Ethereum public networks

	// Mainnet Ethereum mainnet for production
	Mainnet Network = Network{"api", "etherscan.io"}
	// Ropsten Testnet(POW)
	Ropsten Network = Network{"api-ropsten", "etherscan.io"}
	// Kovan Testnet(POA)
	Kovan Network = Network{"api-kovan", "etherscan.io"}
	// Rinkby Testnet(CLIQUE)
	Rinkby Network = Network{"api-rinkeyb", "etherscan.io"}
	// Goerli Testnet(CLIQUE)
	Goerli Network = Network{"api-goerli", "etherscan.io"}
	// Tobalaba Testnet
	Tobalaba Network = Network{"api-tobalaba", "etherscan.io"}
	// Arbitrum Sepolia Testnet
	ArbitrumSepolia Network = Network{"api-sepolia", "arbiscan.io"}
)

// Network is ethereum network type (mainnet, ropsten, etc)
type Network struct {
	subdomain string
	domain    string
}

// SubDomain returns the subdomain of  etherscan API
// via n provided.
func (n Network) SubDomain() (sub string) {
	return n.subdomain
}

func (n Network) Domain() string {
	return n.domain
}
