/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextCreateData struct {
	SupportedFeatures          string            `json:"supportedFeatures,omitempty"`
	UeContext                  UeContext         `json:"ueContext"`
	TargetId                   NgRanTargetId     `json:"targetId"`
	SourceToTargetData         N2InfoContent     `json:"sourceToTargetData"`
	PduSessionList             []N2SmInformation `json:"pduSessionList"`
	UeRadioCapability          *N2InfoContent    `json:"ueRadioCapability,omitempty"`
	UeRadioCapabilityForPaging *N2InfoContent    `json:"ueRadioCapabilityForPaging,omitempty"`
	N2NotifyUri                string            `json:"n2NotifyUri,omitempty"`
	NgapCause                  *NgApCause        `json:"ngapCause,omitempty"`
	ServingNetwork             *PlmnIdNid        `json:"servingNetwork,omitempty"`
}
