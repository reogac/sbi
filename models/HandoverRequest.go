/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverRequest struct {
	AmfUeId               int64                     `json:"amfUeId"`
	SourceToTargetContent []byte                    `json:"sourceToTargetContent,omitempty"`
	HandoverType          int16                     `json:"handoverType"`
	Guami                 Guami                     `json:"guami"`
	UeAmbr                UeAmbr                    `json:"ueAmbr"`
	Sessions              []N2SmInfoDownlinkContent `json:"sessions,omitempty"`
	SecurityContext       SecurityContext           `json:"securityContext"`
	AllowedNssai          AllowedNssai              `json:"allowedNssai"`
	Cause                 N2Cause                   `json:"cause"`
	MaskedImeisv          []byte                    `json:"maskedImeisv,omitempty"`
	NewSecInd             bool                      `json:"newSecInd"`
	UeSecurityCapability  UeSecurityCapability      `json:"ueSecurityCapability"`
}
