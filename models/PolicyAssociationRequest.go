/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:38 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationRequest struct {
	GroupIds          []string      `json:"groupIds,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	Supi              string        `json:"supi"`
	Gpsi              string        `json:"gpsi,omitempty"`
	TimeZone          string        `json:"timeZone,omitempty"`
	NotificationUri   string        `json:"notificationUri"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	Pei               string        `json:"pei,omitempty"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
}
