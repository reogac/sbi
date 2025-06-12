/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:34 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociationRequest struct {
	NotificationUri   string        `json:"notificationUri"`
	AccessType        AccessType    `json:"accessType,omitempty"`
	Pc5Capab          Pc5Capability `json:"pc5Capab,omitempty"`
	ProSeCapab        []string      `json:"proSeCapab,omitempty"`
	Supi              string        `json:"supi"`
	GroupIds          []string      `json:"groupIds,omitempty"`
	UePolReq          string        `json:"uePolReq,omitempty"`
	ServingNfId       string        `json:"servingNfId,omitempty"`
	SuppFeat          string        `json:"suppFeat"`
	AltNotifIpv4Addrs []string      `json:"altNotifIpv4Addrs,omitempty"`
	Gpsi              string        `json:"gpsi,omitempty"`
	TimeZone          string        `json:"timeZone,omitempty"`
	ServingPlmn       *PlmnIdNid    `json:"servingPlmn,omitempty"`
	HPcfId            string        `json:"hPcfId,omitempty"`
	Guami             *Guami        `json:"guami,omitempty"`
	ServiceName       ServiceName   `json:"serviceName,omitempty"`
	AltNotifIpv6Addrs []string      `json:"altNotifIpv6Addrs,omitempty"`
	AltNotifFqdns     []string      `json:"altNotifFqdns,omitempty"`
	Pei               string        `json:"pei,omitempty"`
	UserLoc           *UserLocation `json:"userLoc,omitempty"`
	RatType           RatType       `json:"ratType,omitempty"`
}
