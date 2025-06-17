/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverRequest struct {
	AmfUeId               int64                     `json:"amfUeId"`
	Sessions              []N2SmInfoDownlinkContent `json:"sessions,omitempty"`
	AllowedNssai          AllowedNssai              `json:"allowedNssai"`
	MaskedImeisv          []byte                    `json:"maskedImeisv,omitempty"`
	UeAmbr                UeAmbr                    `json:"ueAmbr"`
	SourceToTargetContent []byte                    `json:"sourceToTargetContent,omitempty"`
	HandoverType          int16                     `json:"handoverType"`
	SecurityContext       SecurityContext           `json:"securityContext"`
	Guami                 Guami                     `json:"guami"`
	Cause                 N2Cause                   `json:"cause"`
	NewSecInd             bool                      `json:"newSecInd"`
	UeSecurityCapability  UeSecurityCapability      `json:"ueSecurityCapability"`
}
