/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextRelocateData struct {
	NgapCause                *NgApCause        `json:"ngapCause,omitempty"`
	SupportedFeatures        string            `json:"supportedFeatures,omitempty"`
	UeContext                UeContext         `json:"ueContext"`
	TargetId                 NgRanTargetId     `json:"targetId"`
	SourceToTargetData       N2InfoContent     `json:"sourceToTargetData"`
	ForwardRelocationRequest RefToBinaryData   `json:"forwardRelocationRequest"`
	PduSessionList           []N2SmInformation `json:"pduSessionList,omitempty"`
	UeRadioCapability        *N2InfoContent    `json:"ueRadioCapability,omitempty"`
}
