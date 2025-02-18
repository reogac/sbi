/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Feb 18 15:05:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverRequest struct {
	SourceToTargetContent []byte                    `json:"sourceToTargetContent,omitempty"`
	HandoverType          int16                     `json:"handoverType"`
	Sessions              []N2SmInfoDownlinkContent `json:"sessions,omitempty"`
	SecurityContext       SecurityContext           `json:"securityContext"`
	Guami                 Guami                     `json:"guami"`
	UeAmbr                UeAmbr                    `json:"ueAmbr"`
	AmfUeId               int64                     `json:"amfUeId"`
	AllowedNssai          AllowedNssai              `json:"allowedNssai"`
	Cause                 N2Cause                   `json:"cause"`
	MaskedImeisv          []byte                    `json:"maskedImeisv,omitempty"`
	NewSecInd             bool                      `json:"newSecInd"`
	UeSecurityCapability  UeSecurityCapability      `json:"ueSecurityCapability"`
}
