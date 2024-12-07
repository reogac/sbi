/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:32 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationRequest struct {
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	Gpsi              string        `json:"gpsi,omitempty"`
	Pei               string        `json:"pei,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	NotificationUri   string        `json:"notificationUri"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	Supi              string        `json:"supi"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	TimeZone          string        `json:"timeZone,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
}
