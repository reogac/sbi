/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:44 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationRequest struct {
	Supi              string        `json:"supi"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	TimeZone          string        `json:"timeZone,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	Pei               string        `json:"pei,omitempty"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	NotificationUri   string        `json:"notificationUri"`
	Gpsi              string        `json:"gpsi,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
}
