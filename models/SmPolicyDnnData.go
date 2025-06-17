/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDnnData struct {
	Ipv6Index           *int                              `json:"ipv6Index,omitempty"`
	ChfInfo             *ChargingInformation              `json:"chfInfo,omitempty"`
	ImsSignallingPrio   *bool                             `json:"imsSignallingPrio,omitempty"`
	PraInfos            map[string]PresenceInfo           `json:"praInfos,omitempty"`
	BdtRefIds           map[string]string                 `json:"bdtRefIds,omitempty"`
	RefUmDataLimitIds   map[string]LimitIdToMonitoringKey `json:"refUmDataLimitIds,omitempty"`
	McsPriority         *bool                             `json:"mcsPriority,omitempty"`
	MpsPriorityLevel    *int                              `json:"mpsPriorityLevel,omitempty"`
	AllowedServices     []string                          `json:"allowedServices,omitempty"`
	GbrUl               string                            `json:"gbrUl,omitempty"`
	GbrDl               string                            `json:"gbrDl,omitempty"`
	SubscSpendingLimits *bool                             `json:"subscSpendingLimits,omitempty"`
	Ipv4Index           *int                              `json:"ipv4Index,omitempty"`
	SubscCats           []string                          `json:"subscCats,omitempty"`
	Offline             *bool                             `json:"offline,omitempty"`
	Online              *bool                             `json:"online,omitempty"`
	McsPriorityLevel    *int                              `json:"mcsPriorityLevel,omitempty"`
	Dnn                 string                            `json:"dnn"`
	AdcSupport          *bool                             `json:"adcSupport,omitempty"`
	MpsPriority         *bool                             `json:"mpsPriority,omitempty"`
	LocRoutNotAllowed   *bool                             `json:"locRoutNotAllowed,omitempty"`
}
