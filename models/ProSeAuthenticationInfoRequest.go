/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProSeAuthenticationInfoRequest struct {
	ServingNetworkName    string                 `json:"servingNetworkName"`
	RelayServiceCode      int                    `json:"relayServiceCode"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
}
