/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SorInfo struct {
	AckInd                  bool               `json:"ackInd"`
	SorMacIausf             string             `json:"sorMacIausf,omitempty"`
	ProvisioningTime        string             `json:"provisioningTime"`
	SorTransparentContainer string             `json:"sorTransparentContainer,omitempty"`
	SteeringContainer       *SteeringContainer `json:"steeringContainer,omitempty"`
	SorCmci                 string             `json:"sorCmci,omitempty"`
	StoreSorCmciInMe        *bool              `json:"storeSorCmciInMe,omitempty"`
	UsimSupportOfSorCmci    *bool              `json:"usimSupportOfSorCmci,omitempty"`
	Countersor              string             `json:"countersor,omitempty"`
}
