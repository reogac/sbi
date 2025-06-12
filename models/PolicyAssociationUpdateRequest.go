/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:34 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationUpdateRequest struct {
	NotificationUri     string                               `json:"notificationUri,omitempty"`
	AltNotifFqdns       []string                             `json:"altNotifFqdns,omitempty"`
	UserLoc             *UserLocation                        `json:"userLoc,omitempty"`
	ConnectState        CmState                              `json:"connectState,omitempty"`
	ProSeCapab          []string                             `json:"proSeCapab,omitempty"`
	AltNotifIpv6Addrs   []string                             `json:"altNotifIpv6Addrs,omitempty"`
	PraStatuses         map[string]PresenceInfo              `json:"praStatuses,omitempty"`
	UePolDelResult      string                               `json:"uePolDelResult,omitempty"`
	UePolReq            string                               `json:"uePolReq,omitempty"`
	PlmnId              *PlmnIdNid                           `json:"plmnId,omitempty"`
	GroupIds            []string                             `json:"groupIds,omitempty"`
	Triggers            []string                             `json:"triggers,omitempty"`
	UePolTransFailNotif *UePolicyTransferFailureNotification `json:"uePolTransFailNotif,omitempty"`
	Guami               *Guami                               `json:"guami,omitempty"`
	AltNotifIpv4Addrs   []string                             `json:"altNotifIpv4Addrs,omitempty"`
	ServingNfId         string                               `json:"servingNfId,omitempty"`
}
