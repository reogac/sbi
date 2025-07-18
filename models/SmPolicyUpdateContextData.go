/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyUpdateContextData struct {
	RelUeMac                 string                       `json:"relUeMac,omitempty"`
	UserLocationInfoTime     string                       `json:"userLocationInfoTime,omitempty"`
	RefQosIndication         *bool                        `json:"refQosIndication,omitempty"`
	TraceReq                 *TraceData                   `json:"traceReq,omitempty"`
	SubsDefQos               *SubscribedDefaultQos        `json:"subsDefQos,omitempty"`
	UserLocationInfo         *UserLocation                `json:"userLocationInfo,omitempty"`
	TypesOfNotif             []string                     `json:"typesOfNotif,omitempty"`
	AccessType               AccessType                   `json:"accessType,omitempty"`
	Ipv6AddressPrefix        string                       `json:"ipv6AddressPrefix,omitempty"`
	TsnBridgeInfo            *TsnBridgeInfo               `json:"tsnBridgeInfo,omitempty"`
	VplmnQosNotApp           *bool                        `json:"vplmnQosNotApp,omitempty"`
	PolicyDecFailureReports  []string                     `json:"policyDecFailureReports,omitempty"`
	NwdafDatas               []NwdafData                  `json:"nwdafDatas,omitempty"`
	AnGwStatus               *bool                        `json:"anGwStatus,omitempty"`
	AddAccessInfo            *AdditionalAccessInfo        `json:"addAccessInfo,omitempty"`
	ThreeGppPsDataOffStatus  *bool                        `json:"3gppPsDataOffStatus,omitempty"`
	InterGrpIds              []string                     `json:"interGrpIds,omitempty"`
	SatBackhaulCategory      SatelliteBackhaulCategory    `json:"satBackhaulCategory,omitempty"`
	RelAccessInfo            *AdditionalAccessInfo        `json:"relAccessInfo,omitempty"`
	SessRuleReports          []SessionRuleReport          `json:"sessRuleReports,omitempty"`
	RepPraInfos              map[string]PresenceInfo      `json:"repPraInfos,omitempty"`
	MaPduInd                 MaPduIndication              `json:"maPduInd,omitempty"`
	TsnPortManContDstt       *PortManagementContainer     `json:"tsnPortManContDstt,omitempty"`
	Ipv4Address              string                       `json:"ipv4Address,omitempty"`
	AuthProfIndex            string                       `json:"authProfIndex,omitempty"`
	UeInitResReq             *UeInitiatedResourceRequest  `json:"ueInitResReq,omitempty"`
	UeTimeZone               string                       `json:"ueTimeZone,omitempty"`
	AddIpv6AddrPrefixes      string                       `json:"addIpv6AddrPrefixes,omitempty"`
	NumOfPackFilter          *int                         `json:"numOfPackFilter,omitempty"`
	AccuUsageReports         []AccuUsageReport            `json:"accuUsageReports,omitempty"`
	TsnBridgeManCont         *BridgeManagementContainer   `json:"tsnBridgeManCont,omitempty"`
	RepPolicyCtrlReqTriggers []string                     `json:"repPolicyCtrlReqTriggers,omitempty"`
	IpDomain                 string                       `json:"ipDomain,omitempty"`
	AppDetectionInfos        []AppDetectionInfo           `json:"appDetectionInfos,omitempty"`
	MulAddrInfos             []IpMulticastAddressInfo     `json:"mulAddrInfos,omitempty"`
	TrafficDescriptors       []DddTrafficDescriptor       `json:"trafficDescriptors,omitempty"`
	RelIpv4Address           string                       `json:"relIpv4Address,omitempty"`
	AddRelIpv6AddrPrefixes   string                       `json:"addRelIpv6AddrPrefixes,omitempty"`
	CreditManageStatus       CreditManagementStatus       `json:"creditManageStatus,omitempty"`
	AtsssCapab               AtsssCapability              `json:"atsssCapab,omitempty"`
	InvalidPolicyDecs        []InvalidParam               `json:"invalidPolicyDecs,omitempty"`
	UeMac                    string                       `json:"ueMac,omitempty"`
	QosMonReports            []QosMonitoringReport        `json:"qosMonReports,omitempty"`
	QosFlowUsage             QosFlowUsage                 `json:"qosFlowUsage,omitempty"`
	PcfUeInfo                *PcfUeCallbackInfo           `json:"pcfUeInfo,omitempty"`
	RuleReports              []RuleReport                 `json:"ruleReports,omitempty"`
	AccNetChIds              []AccNetChId                 `json:"accNetChIds,omitempty"`
	RatType                  RatType                      `json:"ratType,omitempty"`
	ServingNetwork           *PlmnIdNid                   `json:"servingNetwork,omitempty"`
	SubsSessAmbr             *Ambr                        `json:"subsSessAmbr,omitempty"`
	VplmnQos                 *VplmnQos                    `json:"vplmnQos,omitempty"`
	TsnPortManContNwtts      []PortManagementContainer    `json:"tsnPortManContNwtts,omitempty"`
	QncReports               []QosNotificationControlInfo `json:"qncReports,omitempty"`
	PccRuleId                string                       `json:"pccRuleId,omitempty"`
	RelIpv6AddressPrefix     string                       `json:"relIpv6AddressPrefix,omitempty"`
	ServNfId                 *ServingNfIdentity           `json:"servNfId,omitempty"`
}
