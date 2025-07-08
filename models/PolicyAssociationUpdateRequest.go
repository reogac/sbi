/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationUpdateRequest struct {
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
}
