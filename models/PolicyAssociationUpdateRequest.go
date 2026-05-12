/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:44 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationUpdateRequest struct {
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
}
