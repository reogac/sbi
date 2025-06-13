/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:28 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationRequest struct {
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	Supi              string        `json:"supi"`
	Pei               string        `json:"pei,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	Gpsi              string        `json:"gpsi,omitempty"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	TimeZone          string        `json:"timeZone,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	NotificationUri   string        `json:"notificationUri"`
	RatType           RatType       `json:"ratType,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
}
