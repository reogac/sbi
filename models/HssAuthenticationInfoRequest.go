/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:22 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HssAuthenticationInfoRequest struct {
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	AnId                  AccessNetworkId        `json:"anId,omitempty"`
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	HssAuthType           HssAuthType            `json:"hssAuthType"`
	NumOfRequestedVectors int                    `json:"numOfRequestedVectors"`
	RequestingNodeType    NodeType               `json:"requestingNodeType,omitempty"`
	ServingNetworkId      *PlmnId                `json:"servingNetworkId,omitempty"`
}
