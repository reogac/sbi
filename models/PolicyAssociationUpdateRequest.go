/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:01 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationUpdateRequest struct {
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
}
