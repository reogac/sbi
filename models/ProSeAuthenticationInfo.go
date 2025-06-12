/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProSeAuthenticationInfo struct {
	ServingNetworkName string `json:"servingNetworkName"`
	SupportedFeatures  string `json:"supportedFeatures,omitempty"`
	SupiOrSuci         string `json:"supiOrSuci,omitempty"`
	FiveGPrukId        string `json:"5gPrukId,omitempty"`
	RelayServiceCode   int    `json:"relayServiceCode"`
	Nonce1             string `json:"nonce1"`
}
