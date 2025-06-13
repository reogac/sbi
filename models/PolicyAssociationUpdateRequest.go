/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:50 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationUpdateRequest struct {
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
}
