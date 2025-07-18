/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationRequest struct {
	HPcfId            string        `json:"hPcfId,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	TimeZone          string        `json:"timeZone,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	Gpsi              string        `json:"gpsi,omitempty"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	NotificationUri   string        `json:"notificationUri"`
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	Supi              string        `json:"supi"`
	Pei               string        `json:"pei,omitempty"`
}
