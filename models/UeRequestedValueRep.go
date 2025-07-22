/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeRequestedValueRep struct {
	PlmnId       *PlmnIdNid              `json:"plmnId,omitempty"`
	ConnectState CmState                 `json:"connectState,omitempty"`
	UserLoc      *UserLocation           `json:"userLoc,omitempty"`
	PraStatuses  map[string]PresenceInfo `json:"praStatuses,omitempty"`
}
