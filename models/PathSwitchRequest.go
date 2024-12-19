/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:44:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PathSwitchRequest struct {
	RanUeId              RanUeId                 `json:"ranUeId"`
	Loc                  UserLocation            `json:"loc"`
	UeSecurityCapability *UeSecurityCapability   `json:"ueSecurityCapability,omitempty"`
	Sessions             []N2SmInfoUplinkContent `json:"sessions"`
}
