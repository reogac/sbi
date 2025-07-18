/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationUpdateRequest struct {
	Guami               *Guami                               `json:"guami,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
}
