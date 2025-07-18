/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PathSwitchAcknowledge struct {
	Sessions             []N2SmInfoDownlinkContent `json:"sessions"`
	UeSecurityCapability UeSecurityCapability      `json:"ueSecurityCapability"`
	NewSecInd            *bool                     `json:"newSecInd,omitempty"`
	SecurityContext      SecurityContext           `json:"securityContext"`
	AllowedNssai         AllowedNssai              `json:"allowedNssai"`
}
