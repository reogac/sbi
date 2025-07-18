/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextCreateData struct {
	TargetId                   NgRanTargetId     `json:"targetId"`
	PduSessionList             []N2SmInformation `json:"pduSessionList"`
	N2NotifyUri                string            `json:"n2NotifyUri,omitempty"`
	SupportedFeatures          string            `json:"supportedFeatures,omitempty"`
	ServingNetwork             *PlmnIdNid        `json:"servingNetwork,omitempty"`
	UeContext                  UeContext         `json:"ueContext"`
	SourceToTargetData         N2InfoContent     `json:"sourceToTargetData"`
	UeRadioCapability          *N2InfoContent    `json:"ueRadioCapability,omitempty"`
	UeRadioCapabilityForPaging *N2InfoContent    `json:"ueRadioCapabilityForPaging,omitempty"`
	NgapCause                  *NgApCause        `json:"ngapCause,omitempty"`
}
