/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:49:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SorInfo struct {
	SteeringContainer       *SteeringContainer `json:"steeringContainer,omitempty"`
	AckInd                  bool               `json:"ackInd"`
	SorMacIausf             string             `json:"sorMacIausf,omitempty"`
	SorTransparentContainer string             `json:"sorTransparentContainer,omitempty"`
	StoreSorCmciInMe        *bool              `json:"storeSorCmciInMe,omitempty"`
	Countersor              string             `json:"countersor,omitempty"`
	ProvisioningTime        string             `json:"provisioningTime"`
	SorCmci                 string             `json:"sorCmci,omitempty"`
	UsimSupportOfSorCmci    *bool              `json:"usimSupportOfSorCmci,omitempty"`
}
