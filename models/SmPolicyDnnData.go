/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDnnData struct {
	Dnn                 string                            `json:"dnn"`
	SubscSpendingLimits *bool                             `json:"subscSpendingLimits,omitempty"`
	MpsPriority         *bool                             `json:"mpsPriority,omitempty"`
	MpsPriorityLevel    *int                              `json:"mpsPriorityLevel,omitempty"`
	AllowedServices     []string                          `json:"allowedServices,omitempty"`
	GbrUl               string                            `json:"gbrUl,omitempty"`
	GbrDl               string                            `json:"gbrDl,omitempty"`
	AdcSupport          *bool                             `json:"adcSupport,omitempty"`
	Ipv6Index           *int                              `json:"ipv6Index,omitempty"`
	RefUmDataLimitIds   map[string]LimitIdToMonitoringKey `json:"refUmDataLimitIds,omitempty"`
	McsPriority         *bool                             `json:"mcsPriority,omitempty"`
	LocRoutNotAllowed   *bool                             `json:"locRoutNotAllowed,omitempty"`
	SubscCats           []string                          `json:"subscCats,omitempty"`
	Ipv4Index           *int                              `json:"ipv4Index,omitempty"`
	Offline             *bool                             `json:"offline,omitempty"`
	Online              *bool                             `json:"online,omitempty"`
	ChfInfo             *ChargingInformation              `json:"chfInfo,omitempty"`
	PraInfos            map[string]PresenceInfo           `json:"praInfos,omitempty"`
	BdtRefIds           map[string]string                 `json:"bdtRefIds,omitempty"`
	ImsSignallingPrio   *bool                             `json:"imsSignallingPrio,omitempty"`
	McsPriorityLevel    *int                              `json:"mcsPriorityLevel,omitempty"`
}
