/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
