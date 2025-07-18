/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverRequest struct {
	SourceToTargetContent []byte                    `json:"sourceToTargetContent,omitempty"`
	Sessions              []N2SmInfoDownlinkContent `json:"sessions,omitempty"`
	SecurityContext       SecurityContext           `json:"securityContext"`
	Guami                 Guami                     `json:"guami"`
	NewSecInd             bool                      `json:"newSecInd"`
	UeAmbr                UeAmbr                    `json:"ueAmbr"`
	AmfUeId               int64                     `json:"amfUeId"`
	HandoverType          int16                     `json:"handoverType"`
	AllowedNssai          AllowedNssai              `json:"allowedNssai"`
	Cause                 N2Cause                   `json:"cause"`
	MaskedImeisv          []byte                    `json:"maskedImeisv,omitempty"`
	UeSecurityCapability  UeSecurityCapability      `json:"ueSecurityCapability"`
}
