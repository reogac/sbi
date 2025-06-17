/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:01 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationRequest struct {
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	Supi              string        `json:"supi"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	NotificationUri   string        `json:"notificationUri"`
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	Gpsi              string        `json:"gpsi,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	Pei               string        `json:"pei,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	TimeZone          string        `json:"timeZone,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
}
