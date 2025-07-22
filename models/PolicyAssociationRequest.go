/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationRequest struct {
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	Pei               string        `json:"pei,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	TimeZone          string        `json:"timeZone,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	NotificationUri   string        `json:"notificationUri"`
	Supi              string        `json:"supi"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	Gpsi              string        `json:"gpsi,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
}
