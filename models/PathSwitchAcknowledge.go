/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 11:07:25 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PathSwitchAcknowledge struct {
	Sessions             []N2SmInfoDownlinkContent `json:"sessions"`
	UeSecurityCapability UeSecurityCapability      `json:"ueSecurityCapability"`
	NewSecInt            *bool                     `json:"newSecInt,omitempty"`
	SecurityContext      SecurityContext           `json:"securityContext"`
	AllowedNssai         AllowedNssai              `json:"allowedNssai"`
}
