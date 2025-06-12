/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverRequest struct {
	AmfUeId               int64                     `json:"amfUeId"`
	NewSecInd             bool                      `json:"newSecInd"`
	SecurityContext       SecurityContext           `json:"securityContext"`
	Guami                 Guami                     `json:"guami"`
	AllowedNssai          AllowedNssai              `json:"allowedNssai"`
	Cause                 N2Cause                   `json:"cause"`
	MaskedImeisv          []byte                    `json:"maskedImeisv,omitempty"`
	SourceToTargetContent []byte                    `json:"sourceToTargetContent,omitempty"`
	HandoverType          int16                     `json:"handoverType"`
	Sessions              []N2SmInfoDownlinkContent `json:"sessions,omitempty"`
	UeSecurityCapability  UeSecurityCapability      `json:"ueSecurityCapability"`
	UeAmbr                UeAmbr                    `json:"ueAmbr"`
}
