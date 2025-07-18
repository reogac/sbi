/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProSeAuthenticationInfo struct {
	SupiOrSuci         string `json:"supiOrSuci,omitempty"`
	FiveGPrukId        string `json:"5gPrukId,omitempty"`
	RelayServiceCode   int    `json:"relayServiceCode"`
	Nonce1             string `json:"nonce1"`
	ServingNetworkName string `json:"servingNetworkName"`
	SupportedFeatures  string `json:"supportedFeatures,omitempty"`
}
