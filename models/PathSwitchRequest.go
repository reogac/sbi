/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PathSwitchRequest struct {
	RanUeId          RanUeId                 `json:"ranUeId"`
	Loc              *UserLocation           `json:"loc,omitempty"`
	UeSecCap         *UeSecurityCapability   `json:"ueSecCap,omitempty"`
	SwitchedSessions []N2SmInfoUplinkContent `json:"switchedSessions,omitempty"`
	FailedSessions   []N2SmInfoUplinkContent `json:"failedSessions,omitempty"`
}
