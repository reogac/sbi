/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDnnData struct {
	AllowedServices     []string                          `json:"allowedServices,omitempty"`
	GbrUl               string                            `json:"gbrUl,omitempty"`
	GbrDl               string                            `json:"gbrDl,omitempty"`
	Ipv4Index           *int                              `json:"ipv4Index,omitempty"`
	Online              *bool                             `json:"online,omitempty"`
	LocRoutNotAllowed   *bool                             `json:"locRoutNotAllowed,omitempty"`
	McsPriorityLevel    *int                              `json:"mcsPriorityLevel,omitempty"`
	SubscCats           []string                          `json:"subscCats,omitempty"`
	AdcSupport          *bool                             `json:"adcSupport,omitempty"`
	SubscSpendingLimits *bool                             `json:"subscSpendingLimits,omitempty"`
	Ipv6Index           *int                              `json:"ipv6Index,omitempty"`
	Offline             *bool                             `json:"offline,omitempty"`
	ChfInfo             *ChargingInformation              `json:"chfInfo,omitempty"`
	ImsSignallingPrio   *bool                             `json:"imsSignallingPrio,omitempty"`
	BdtRefIds           map[string]string                 `json:"bdtRefIds,omitempty"`
	Dnn                 string                            `json:"dnn"`
	RefUmDataLimitIds   map[string]LimitIdToMonitoringKey `json:"refUmDataLimitIds,omitempty"`
	PraInfos            map[string]PresenceInfo           `json:"praInfos,omitempty"`
	MpsPriority         *bool                             `json:"mpsPriority,omitempty"`
	McsPriority         *bool                             `json:"mcsPriority,omitempty"`
	MpsPriorityLevel    *int                              `json:"mpsPriorityLevel,omitempty"`
}
