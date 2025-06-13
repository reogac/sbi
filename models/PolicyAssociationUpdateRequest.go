/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:28 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationUpdateRequest struct {
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
}
