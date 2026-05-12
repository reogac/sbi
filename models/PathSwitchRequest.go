/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:27 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PathSwitchRequest struct {
	Loc                  UserLocation            `json:"loc"`
	UeSecurityCapability *UeSecurityCapability   `json:"ueSecurityCapability,omitempty"`
	Sessions             []N2SmInfoUplinkContent `json:"sessions"`
	RanNets              []string                `json:"ranNets,omitempty"`
	RanUeId              RanUeId                 `json:"ranUeId"`
}
