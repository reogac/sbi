/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:30 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDecision struct {
	RelCause              SmPolicyAssociationReleaseCause `json:"relCause,omitempty"`
	PcscfRestIndication   *bool                           `json:"pcscfRestIndication,omitempty"`
	Conds                 map[string]*ConditionData       `json:"conds,omitempty"`
	LastReqUsageData      *RequestedUsageData             `json:"lastReqUsageData,omitempty"`
	TsnPortManContDstt    *PortManagementContainer        `json:"tsnPortManContDstt,omitempty"`
	ChargingInfo          *ChargingInformation            `json:"chargingInfo,omitempty"`
	SuppFeat              string                          `json:"suppFeat,omitempty"`
	ReflectiveQoSTimer    *int                            `json:"reflectiveQoSTimer,omitempty"`
	Offline               *bool                           `json:"offline,omitempty"`
	OfflineChOnly         *bool                           `json:"offlineChOnly,omitempty"`
	PolicyCtrlReqTriggers []string                        `json:"policyCtrlReqTriggers,omitempty"`
	QosFlowUsage          QosFlowUsage                    `json:"qosFlowUsage,omitempty"`
	UmDecs                map[string]*UsageMonitoringData `json:"umDecs,omitempty"`
	ChgDecs               map[string]*ChargingData        `json:"chgDecs,omitempty"`
	PraInfos              map[string]*PresenceInfoRm      `json:"praInfos,omitempty"`
	TsnBridgeManCont      *BridgeManagementContainer      `json:"tsnBridgeManCont,omitempty"`
	QosDecs               map[string]*QosData             `json:"qosDecs,omitempty"`
	QosMonDecs            map[string]*QosMonitoringData   `json:"qosMonDecs,omitempty"`
	LastReqRuleData       []RequestedRuleData             `json:"lastReqRuleData,omitempty"`
	TsnPortManContNwtts   []PortManagementContainer       `json:"tsnPortManContNwtts,omitempty"`
	SessRules             map[string]*SessionRule         `json:"sessRules,omitempty"`
	TraffContDecs         map[string]*TrafficControlData  `json:"traffContDecs,omitempty"`
	Online                *bool                           `json:"online,omitempty"`
	Ipv4Index             *int                            `json:"ipv4Index,omitempty"`
	Ipv6Index             *int                            `json:"ipv6Index,omitempty"`
	PccRules              map[string]*PccRule             `json:"pccRules,omitempty"`
	RevalidationTime      string                          `json:"revalidationTime,omitempty"`
	RedSessIndication     *bool                           `json:"redSessIndication,omitempty"`
	QosChars              map[string]*QosCharacteristics  `json:"qosChars,omitempty"`
}
