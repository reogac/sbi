/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:01 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeRequestedValueRep struct {
	UserLoc      *UserLocation           `json:"userLoc,omitempty"`
	PraStatuses  map[string]PresenceInfo `json:"praStatuses,omitempty"`
	PlmnId       *PlmnIdNid              `json:"plmnId,omitempty"`
	ConnectState CmState                 `json:"connectState,omitempty"`
}
