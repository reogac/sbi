/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:38 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProSeAuthenticationInfoRequest struct {
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	ServingNetworkName    string                 `json:"servingNetworkName"`
	RelayServiceCode      int                    `json:"relayServiceCode"`
}
