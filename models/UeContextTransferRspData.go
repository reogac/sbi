/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextTransferRspData struct {
	UeContext                  UeContext      `json:"ueContext"`
	UeRadioCapability          *N2InfoContent `json:"ueRadioCapability,omitempty"`
	UeRadioCapabilityForPaging *N2InfoContent `json:"ueRadioCapabilityForPaging,omitempty"`
	UeNbiotRadioCapability     *N2InfoContent `json:"ueNbiotRadioCapability,omitempty"`
	SupportedFeatures          string         `json:"supportedFeatures,omitempty"`
}
