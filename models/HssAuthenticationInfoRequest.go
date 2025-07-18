/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:42 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HssAuthenticationInfoRequest struct {
	SupportedFeatures     string                 `json:"supportedFeatures,omitempty"`
	HssAuthType           HssAuthType            `json:"hssAuthType"`
	NumOfRequestedVectors int                    `json:"numOfRequestedVectors"`
	RequestingNodeType    NodeType               `json:"requestingNodeType,omitempty"`
	ServingNetworkId      *PlmnId                `json:"servingNetworkId,omitempty"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	AnId                  AccessNetworkId        `json:"anId,omitempty"`
}
