/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextCreateData struct {
	NgapCause                  *NgApCause        `json:"ngapCause,omitempty"`
	SupportedFeatures          string            `json:"supportedFeatures,omitempty"`
	TargetId                   NgRanTargetId     `json:"targetId"`
	PduSessionList             []N2SmInformation `json:"pduSessionList"`
	N2NotifyUri                string            `json:"n2NotifyUri,omitempty"`
	UeRadioCapability          *N2InfoContent    `json:"ueRadioCapability,omitempty"`
	UeContext                  UeContext         `json:"ueContext"`
	SourceToTargetData         N2InfoContent     `json:"sourceToTargetData"`
	UeRadioCapabilityForPaging *N2InfoContent    `json:"ueRadioCapabilityForPaging,omitempty"`
	ServingNetwork             *PlmnIdNid        `json:"servingNetwork,omitempty"`
}
