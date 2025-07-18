/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:38 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationUpdateRequest struct {
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
}
