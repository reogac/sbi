/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:50 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationRequest struct {
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	NotificationUri   string        `json:"notificationUri"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	Supi              string        `json:"supi"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	Gpsi              string        `json:"gpsi,omitempty"`
	Pei               string        `json:"pei,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	TimeZone          string        `json:"timeZone,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
}
