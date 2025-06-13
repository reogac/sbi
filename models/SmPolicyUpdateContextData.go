/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyUpdateContextData struct {
	PolicyDecFailureReports  []string                     `json:"policyDecFailureReports,omitempty"`
	AppDetectionInfos        []AppDetectionInfo           `json:"appDetectionInfos,omitempty"`
	RuleReports              []RuleReport                 `json:"ruleReports,omitempty"`
	SessRuleReports          []SessionRuleReport          `json:"sessRuleReports,omitempty"`
	IpDomain                 string                       `json:"ipDomain,omitempty"`
	QncReports               []QosNotificationControlInfo `json:"qncReports,omitempty"`
	TsnBridgeManCont         *BridgeManagementContainer   `json:"tsnBridgeManCont,omitempty"`
	AddAccessInfo            *AdditionalAccessInfo        `json:"addAccessInfo,omitempty"`
	VplmnQosNotApp           *bool                        `json:"vplmnQosNotApp,omitempty"`
	AccessType               AccessType                   `json:"accessType,omitempty"`
	NwdafDatas               []NwdafData                  `json:"nwdafDatas,omitempty"`
	RelIpv6AddressPrefix     string                       `json:"relIpv6AddressPrefix,omitempty"`
	UserLocationInfoTime     string                       `json:"userLocationInfoTime,omitempty"`
	QosFlowUsage             QosFlowUsage                 `json:"qosFlowUsage,omitempty"`
	UeInitResReq             *UeInitiatedResourceRequest  `json:"ueInitResReq,omitempty"`
	RelUeMac                 string                       `json:"relUeMac,omitempty"`
	AuthProfIndex            string                       `json:"authProfIndex,omitempty"`
	AtsssCapab               AtsssCapability              `json:"atsssCapab,omitempty"`
	TsnPortManContDstt       *PortManagementContainer     `json:"tsnPortManContDstt,omitempty"`
	TsnBridgeInfo            *TsnBridgeInfo               `json:"tsnBridgeInfo,omitempty"`
	MulAddrInfos             []IpMulticastAddressInfo     `json:"mulAddrInfos,omitempty"`
	InterGrpIds              []string                     `json:"interGrpIds,omitempty"`
	AnGwStatus               *bool                        `json:"anGwStatus,omitempty"`
	ServingNetwork           *PlmnIdNid                   `json:"servingNetwork,omitempty"`
	UeMac                    string                       `json:"ueMac,omitempty"`
	CreditManageStatus       CreditManagementStatus       `json:"creditManageStatus,omitempty"`
	SatBackhaulCategory      SatelliteBackhaulCategory    `json:"satBackhaulCategory,omitempty"`
	RepPolicyCtrlReqTriggers []string                     `json:"repPolicyCtrlReqTriggers,omitempty"`
	UeTimeZone               string                       `json:"ueTimeZone,omitempty"`
	Ipv6AddressPrefix        string                       `json:"ipv6AddressPrefix,omitempty"`
	TrafficDescriptors       []DddTrafficDescriptor       `json:"trafficDescriptors,omitempty"`
	AccNetChIds              []AccNetChId                 `json:"accNetChIds,omitempty"`
	RelIpv4Address           string                       `json:"relIpv4Address,omitempty"`
	NumOfPackFilter          *int                         `json:"numOfPackFilter,omitempty"`
	QosMonReports            []QosMonitoringReport        `json:"qosMonReports,omitempty"`
	ServNfId                 *ServingNfIdentity           `json:"servNfId,omitempty"`
	MaPduInd                 MaPduIndication              `json:"maPduInd,omitempty"`
	PccRuleId                string                       `json:"pccRuleId,omitempty"`
	RatType                  RatType                      `json:"ratType,omitempty"`
	UserLocationInfo         *UserLocation                `json:"userLocationInfo,omitempty"`
	Ipv4Address              string                       `json:"ipv4Address,omitempty"`
	SubsDefQos               *SubscribedDefaultQos        `json:"subsDefQos,omitempty"`
	ThreeGppPsDataOffStatus  *bool                        `json:"3gppPsDataOffStatus,omitempty"`
	RefQosIndication         *bool                        `json:"refQosIndication,omitempty"`
	TraceReq                 *TraceData                   `json:"traceReq,omitempty"`
	RelAccessInfo            *AdditionalAccessInfo        `json:"relAccessInfo,omitempty"`
	TsnPortManContNwtts      []PortManagementContainer    `json:"tsnPortManContNwtts,omitempty"`
	TypesOfNotif             []string                     `json:"typesOfNotif,omitempty"`
	VplmnQos                 *VplmnQos                    `json:"vplmnQos,omitempty"`
	AddIpv6AddrPrefixes      string                       `json:"addIpv6AddrPrefixes,omitempty"`
	AddRelIpv6AddrPrefixes   string                       `json:"addRelIpv6AddrPrefixes,omitempty"`
	SubsSessAmbr             *Ambr                        `json:"subsSessAmbr,omitempty"`
	PcfUeInfo                *PcfUeCallbackInfo           `json:"pcfUeInfo,omitempty"`
	AccuUsageReports         []AccuUsageReport            `json:"accuUsageReports,omitempty"`
	RepPraInfos              map[string]PresenceInfo      `json:"repPraInfos,omitempty"`
	InvalidPolicyDecs        []InvalidParam               `json:"invalidPolicyDecs,omitempty"`
}
