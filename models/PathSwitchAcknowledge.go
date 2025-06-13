/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:14 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PathSwitchAcknowledge struct {
	SecurityContext      SecurityContext           `json:"securityContext"`
	AllowedNssai         AllowedNssai              `json:"allowedNssai"`
	Sessions             []N2SmInfoDownlinkContent `json:"sessions"`
	UeSecurityCapability UeSecurityCapability      `json:"ueSecurityCapability"`
	NewSecInd            *bool                     `json:"newSecInd,omitempty"`
}
