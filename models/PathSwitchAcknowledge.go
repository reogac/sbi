/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:27 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PathSwitchAcknowledge struct {
	AllowedNssai         AllowedNssai              `json:"allowedNssai"`
	Sessions             []N2SmInfoDownlinkContent `json:"sessions"`
	UeSecurityCapability UeSecurityCapability      `json:"ueSecurityCapability"`
	NewSecInd            *bool                     `json:"newSecInd,omitempty"`
	SecurityContext      SecurityContext           `json:"securityContext"`
}
