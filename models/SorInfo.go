/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SorInfo struct {
	Countersor              string             `json:"countersor,omitempty"`
	ProvisioningTime        string             `json:"provisioningTime"`
	SteeringContainer       *SteeringContainer `json:"steeringContainer,omitempty"`
	AckInd                  bool               `json:"ackInd"`
	SorMacIausf             string             `json:"sorMacIausf,omitempty"`
	SorTransparentContainer string             `json:"sorTransparentContainer,omitempty"`
	SorCmci                 string             `json:"sorCmci,omitempty"`
	StoreSorCmciInMe        *bool              `json:"storeSorCmciInMe,omitempty"`
	UsimSupportOfSorCmci    *bool              `json:"usimSupportOfSorCmci,omitempty"`
}
