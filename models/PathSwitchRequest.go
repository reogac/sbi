/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PathSwitchRequest struct {
	UeSecurityCapability *UeSecurityCapability   `json:"ueSecurityCapability,omitempty"`
	Sessions             []N2SmInfoUplinkContent `json:"sessions"`
	Loc                  UserLocation            `json:"loc"`
}
