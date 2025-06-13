/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextCreateData struct {
	SupportedFeatures          string            `json:"supportedFeatures,omitempty"`
	ServingNetwork             *PlmnIdNid        `json:"servingNetwork,omitempty"`
	TargetId                   NgRanTargetId     `json:"targetId"`
	N2NotifyUri                string            `json:"n2NotifyUri,omitempty"`
	UeRadioCapabilityForPaging *N2InfoContent    `json:"ueRadioCapabilityForPaging,omitempty"`
	UeRadioCapability          *N2InfoContent    `json:"ueRadioCapability,omitempty"`
	NgapCause                  *NgApCause        `json:"ngapCause,omitempty"`
	UeContext                  UeContext         `json:"ueContext"`
	SourceToTargetData         N2InfoContent     `json:"sourceToTargetData"`
	PduSessionList             []N2SmInformation `json:"pduSessionList"`
}
