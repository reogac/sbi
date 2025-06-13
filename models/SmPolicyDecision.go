/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 09:53:53 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDecision struct {
	TraffContDecs         map[string]*TrafficControlData  `json:"traffContDecs,omitempty"`
	Offline               *bool                           `json:"offline,omitempty"`
	SessRules             map[string]*SessionRule         `json:"sessRules,omitempty"`
	PccRules              map[string]*PccRule             `json:"pccRules,omitempty"`
	TsnPortManContNwtts   []PortManagementContainer       `json:"tsnPortManContNwtts,omitempty"`
	ChgDecs               map[string]*ChargingData        `json:"chgDecs,omitempty"`
	QosMonDecs            map[string]*QosMonitoringData   `json:"qosMonDecs,omitempty"`
	Online                *bool                           `json:"online,omitempty"`
	QosFlowUsage          QosFlowUsage                    `json:"qosFlowUsage,omitempty"`
	PcscfRestIndication   *bool                           `json:"pcscfRestIndication,omitempty"`
	Conds                 map[string]*ConditionData       `json:"conds,omitempty"`
	LastReqRuleData       []RequestedRuleData             `json:"lastReqRuleData,omitempty"`
	Ipv4Index             *int                            `json:"ipv4Index,omitempty"`
	TsnBridgeManCont      *BridgeManagementContainer      `json:"tsnBridgeManCont,omitempty"`
	RedSessIndication     *bool                           `json:"redSessIndication,omitempty"`
	ChargingInfo          *ChargingInformation            `json:"chargingInfo,omitempty"`
	RevalidationTime      string                          `json:"revalidationTime,omitempty"`
	RelCause              SmPolicyAssociationReleaseCause `json:"relCause,omitempty"`
	OfflineChOnly         *bool                           `json:"offlineChOnly,omitempty"`
	Ipv6Index             *int                            `json:"ipv6Index,omitempty"`
	UmDecs                map[string]*UsageMonitoringData `json:"umDecs,omitempty"`
	PolicyCtrlReqTriggers []string                        `json:"policyCtrlReqTriggers,omitempty"`
	PraInfos              map[string]*PresenceInfoRm      `json:"praInfos,omitempty"`
	SuppFeat              string                          `json:"suppFeat,omitempty"`
	TsnPortManContDstt    *PortManagementContainer        `json:"tsnPortManContDstt,omitempty"`
	QosChars              map[string]*QosCharacteristics  `json:"qosChars,omitempty"`
	ReflectiveQoSTimer    *int                            `json:"reflectiveQoSTimer,omitempty"`
	QosDecs               map[string]*QosData             `json:"qosDecs,omitempty"`
	LastReqUsageData      *RequestedUsageData             `json:"lastReqUsageData,omitempty"`
}
