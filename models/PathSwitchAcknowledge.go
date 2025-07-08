/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:28 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PathSwitchAcknowledge struct {
	UeSecurityCapability UeSecurityCapability      `json:"ueSecurityCapability"`
	NewSecInd            *bool                     `json:"newSecInd,omitempty"`
	SecurityContext      SecurityContext           `json:"securityContext"`
	AllowedNssai         AllowedNssai              `json:"allowedNssai"`
	Sessions             []N2SmInfoDownlinkContent `json:"sessions"`
}
