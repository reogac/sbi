/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDecision struct {
	Online                *bool                           `json:"online,omitempty"`
	PccRules              map[string]*PccRule             `json:"pccRules,omitempty"`
	ChgDecs               map[string]*ChargingData        `json:"chgDecs,omitempty"`
	QosChars              map[string]*QosCharacteristics  `json:"qosChars,omitempty"`
	LastReqUsageData      *RequestedUsageData             `json:"lastReqUsageData,omitempty"`
	Ipv4Index             *int                            `json:"ipv4Index,omitempty"`
	TsnPortManContDstt    *PortManagementContainer        `json:"tsnPortManContDstt,omitempty"`
	TsnPortManContNwtts   []PortManagementContainer       `json:"tsnPortManContNwtts,omitempty"`
	TraffContDecs         map[string]*TrafficControlData  `json:"traffContDecs,omitempty"`
	QosMonDecs            map[string]*QosMonitoringData   `json:"qosMonDecs,omitempty"`
	Conds                 map[string]*ConditionData       `json:"conds,omitempty"`
	SuppFeat              string                          `json:"suppFeat,omitempty"`
	RedSessIndication     *bool                           `json:"redSessIndication,omitempty"`
	ReflectiveQoSTimer    *int                            `json:"reflectiveQoSTimer,omitempty"`
	OfflineChOnly         *bool                           `json:"offlineChOnly,omitempty"`
	PolicyCtrlReqTriggers []string                        `json:"policyCtrlReqTriggers,omitempty"`
	QosDecs               map[string]*QosData             `json:"qosDecs,omitempty"`
	RevalidationTime      string                          `json:"revalidationTime,omitempty"`
	PcscfRestIndication   *bool                           `json:"pcscfRestIndication,omitempty"`
	UmDecs                map[string]*UsageMonitoringData `json:"umDecs,omitempty"`
	Ipv6Index             *int                            `json:"ipv6Index,omitempty"`
	TsnBridgeManCont      *BridgeManagementContainer      `json:"tsnBridgeManCont,omitempty"`
	PraInfos              map[string]*PresenceInfoRm      `json:"praInfos,omitempty"`
	QosFlowUsage          QosFlowUsage                    `json:"qosFlowUsage,omitempty"`
	RelCause              SmPolicyAssociationReleaseCause `json:"relCause,omitempty"`
	SessRules             map[string]*SessionRule         `json:"sessRules,omitempty"`
	Offline               *bool                           `json:"offline,omitempty"`
	LastReqRuleData       []RequestedRuleData             `json:"lastReqRuleData,omitempty"`
	ChargingInfo          *ChargingInformation            `json:"chargingInfo,omitempty"`
}
