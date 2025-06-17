/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:42 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
