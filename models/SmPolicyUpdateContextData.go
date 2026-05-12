/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:42 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyUpdateContextData struct {
	Ipv6AddressPrefix        string                       `json:"ipv6AddressPrefix,omitempty"`
	AnGwStatus               *bool                        `json:"anGwStatus,omitempty"`
	UserLocationInfo         *UserLocation                `json:"userLocationInfo,omitempty"`
	UeMac                    string                       `json:"ueMac,omitempty"`
	AuthProfIndex            string                       `json:"authProfIndex,omitempty"`
	QncReports               []QosNotificationControlInfo `json:"qncReports,omitempty"`
	QosFlowUsage             QosFlowUsage                 `json:"qosFlowUsage,omitempty"`
	CreditManageStatus       CreditManagementStatus       `json:"creditManageStatus,omitempty"`
	InvalidPolicyDecs        []InvalidParam               `json:"invalidPolicyDecs,omitempty"`
	TrafficDescriptors       []DddTrafficDescriptor       `json:"trafficDescriptors,omitempty"`
	RelAccessInfo            *AdditionalAccessInfo        `json:"relAccessInfo,omitempty"`
	RepPraInfos              map[string]PresenceInfo      `json:"repPraInfos,omitempty"`
	MulAddrInfos             []IpMulticastAddressInfo     `json:"mulAddrInfos,omitempty"`
	PcfUeInfo                *PcfUeCallbackInfo           `json:"pcfUeInfo,omitempty"`
	UeTimeZone               string                       `json:"ueTimeZone,omitempty"`
	Ipv4Address              string                       `json:"ipv4Address,omitempty"`
	VplmnQos                 *VplmnQos                    `json:"vplmnQos,omitempty"`
	NumOfPackFilter          *int                         `json:"numOfPackFilter,omitempty"`
	QosMonReports            []QosMonitoringReport        `json:"qosMonReports,omitempty"`
	TraceReq                 *TraceData                   `json:"traceReq,omitempty"`
	MaPduInd                 MaPduIndication              `json:"maPduInd,omitempty"`
	TsnPortManContDstt       *PortManagementContainer     `json:"tsnPortManContDstt,omitempty"`
	RelIpv4Address           string                       `json:"relIpv4Address,omitempty"`
	RelIpv6AddressPrefix     string                       `json:"relIpv6AddressPrefix,omitempty"`
	AddIpv6AddrPrefixes      string                       `json:"addIpv6AddrPrefixes,omitempty"`
	AddRelIpv6AddrPrefixes   string                       `json:"addRelIpv6AddrPrefixes,omitempty"`
	RelUeMac                 string                       `json:"relUeMac,omitempty"`
	SubsDefQos               *SubscribedDefaultQos        `json:"subsDefQos,omitempty"`
	SessRuleReports          []SessionRuleReport          `json:"sessRuleReports,omitempty"`
	UserLocationInfoTime     string                       `json:"userLocationInfoTime,omitempty"`
	RepPolicyCtrlReqTriggers []string                     `json:"repPolicyCtrlReqTriggers,omitempty"`
	AccessType               AccessType                   `json:"accessType,omitempty"`
	SubsSessAmbr             *Ambr                        `json:"subsSessAmbr,omitempty"`
	VplmnQosNotApp           *bool                        `json:"vplmnQosNotApp,omitempty"`
	AccuUsageReports         []AccuUsageReport            `json:"accuUsageReports,omitempty"`
	RefQosIndication         *bool                        `json:"refQosIndication,omitempty"`
	TsnBridgeInfo            *TsnBridgeInfo               `json:"tsnBridgeInfo,omitempty"`
	TsnBridgeManCont         *BridgeManagementContainer   `json:"tsnBridgeManCont,omitempty"`
	AddAccessInfo            *AdditionalAccessInfo        `json:"addAccessInfo,omitempty"`
	AppDetectionInfos        []AppDetectionInfo           `json:"appDetectionInfos,omitempty"`
	RuleReports              []RuleReport                 `json:"ruleReports,omitempty"`
	UeInitResReq             *UeInitiatedResourceRequest  `json:"ueInitResReq,omitempty"`
	PccRuleId                string                       `json:"pccRuleId,omitempty"`
	TypesOfNotif             []string                     `json:"typesOfNotif,omitempty"`
	InterGrpIds              []string                     `json:"interGrpIds,omitempty"`
	NwdafDatas               []NwdafData                  `json:"nwdafDatas,omitempty"`
	RatType                  RatType                      `json:"ratType,omitempty"`
	ServNfId                 *ServingNfIdentity           `json:"servNfId,omitempty"`
	AtsssCapab               AtsssCapability              `json:"atsssCapab,omitempty"`
	TsnPortManContNwtts      []PortManagementContainer    `json:"tsnPortManContNwtts,omitempty"`
	PolicyDecFailureReports  []string                     `json:"policyDecFailureReports,omitempty"`
	SatBackhaulCategory      SatelliteBackhaulCategory    `json:"satBackhaulCategory,omitempty"`
	ThreeGppPsDataOffStatus  *bool                        `json:"3gppPsDataOffStatus,omitempty"`
	AccNetChIds              []AccNetChId                 `json:"accNetChIds,omitempty"`
	ServingNetwork           *PlmnIdNid                   `json:"servingNetwork,omitempty"`
	IpDomain                 string                       `json:"ipDomain,omitempty"`
}
