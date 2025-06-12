/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:16 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PathSwitchRequest struct {
	Loc                  UserLocation            `json:"loc"`
	UeSecurityCapability *UeSecurityCapability   `json:"ueSecurityCapability,omitempty"`
	Sessions             []N2SmInfoUplinkContent `json:"sessions"`
}
