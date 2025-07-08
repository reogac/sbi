/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationRequest struct {
	UePolReq          string        `json:"uePolReq,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	Supi              string        `json:"supi"`
	Pei               string        `json:"pei,omitempty"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	NotificationUri   string        `json:"notificationUri"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	Gpsi              string        `json:"gpsi,omitempty"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	TimeZone          string        `json:"timeZone,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
}
