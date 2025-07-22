/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationUpdateRequest struct {
	ConnectState        CmState                              `json:"connectState,omitempty"`
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
}
