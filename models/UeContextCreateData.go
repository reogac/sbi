/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextCreateData struct {
	SourceToTargetData         N2InfoContent     `json:"sourceToTargetData"`
	PduSessionList             []N2SmInformation `json:"pduSessionList"`
	UeRadioCapabilityForPaging *N2InfoContent    `json:"ueRadioCapabilityForPaging,omitempty"`
	ServingNetwork             *PlmnIdNid        `json:"servingNetwork,omitempty"`
	N2NotifyUri                string            `json:"n2NotifyUri,omitempty"`
	UeRadioCapability          *N2InfoContent    `json:"ueRadioCapability,omitempty"`
	NgapCause                  *NgApCause        `json:"ngapCause,omitempty"`
	SupportedFeatures          string            `json:"supportedFeatures,omitempty"`
	UeContext                  UeContext         `json:"ueContext"`
	TargetId                   NgRanTargetId     `json:"targetId"`
}
