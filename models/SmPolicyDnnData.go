/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:52 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDnnData struct {
	GbrUl               string                            `json:"gbrUl,omitempty"`
	Ipv6Index           *int                              `json:"ipv6Index,omitempty"`
	Offline             *bool                             `json:"offline,omitempty"`
	ChfInfo             *ChargingInformation              `json:"chfInfo,omitempty"`
	RefUmDataLimitIds   map[string]LimitIdToMonitoringKey `json:"refUmDataLimitIds,omitempty"`
	MpsPriority         *bool                             `json:"mpsPriority,omitempty"`
	LocRoutNotAllowed   *bool                             `json:"locRoutNotAllowed,omitempty"`
	GbrDl               string                            `json:"gbrDl,omitempty"`
	AdcSupport          *bool                             `json:"adcSupport,omitempty"`
	SubscSpendingLimits *bool                             `json:"subscSpendingLimits,omitempty"`
	Ipv4Index           *int                              `json:"ipv4Index,omitempty"`
	McsPriorityLevel    *int                              `json:"mcsPriorityLevel,omitempty"`
	PraInfos            map[string]PresenceInfo           `json:"praInfos,omitempty"`
	SubscCats           []string                          `json:"subscCats,omitempty"`
	Online              *bool                             `json:"online,omitempty"`
	ImsSignallingPrio   *bool                             `json:"imsSignallingPrio,omitempty"`
	Dnn                 string                            `json:"dnn"`
	AllowedServices     []string                          `json:"allowedServices,omitempty"`
	McsPriority         *bool                             `json:"mcsPriority,omitempty"`
	MpsPriorityLevel    *int                              `json:"mpsPriorityLevel,omitempty"`
	BdtRefIds           map[string]string                 `json:"bdtRefIds,omitempty"`
}
