/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:47 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDnnData struct {
	McsPriority         *bool                             `json:"mcsPriority,omitempty"`
	ImsSignallingPrio   *bool                             `json:"imsSignallingPrio,omitempty"`
	AllowedServices     []string                          `json:"allowedServices,omitempty"`
	SubscCats           []string                          `json:"subscCats,omitempty"`
	RefUmDataLimitIds   map[string]LimitIdToMonitoringKey `json:"refUmDataLimitIds,omitempty"`
	SubscSpendingLimits *bool                             `json:"subscSpendingLimits,omitempty"`
	Ipv4Index           *int                              `json:"ipv4Index,omitempty"`
	ChfInfo             *ChargingInformation              `json:"chfInfo,omitempty"`
	PraInfos            map[string]PresenceInfo           `json:"praInfos,omitempty"`
	GbrUl               string                            `json:"gbrUl,omitempty"`
	GbrDl               string                            `json:"gbrDl,omitempty"`
	AdcSupport          *bool                             `json:"adcSupport,omitempty"`
	BdtRefIds           map[string]string                 `json:"bdtRefIds,omitempty"`
	Ipv6Index           *int                              `json:"ipv6Index,omitempty"`
	Offline             *bool                             `json:"offline,omitempty"`
	MpsPriority         *bool                             `json:"mpsPriority,omitempty"`
	McsPriorityLevel    *int                              `json:"mcsPriorityLevel,omitempty"`
	LocRoutNotAllowed   *bool                             `json:"locRoutNotAllowed,omitempty"`
	Dnn                 string                            `json:"dnn"`
	Online              *bool                             `json:"online,omitempty"`
	MpsPriorityLevel    *int                              `json:"mpsPriorityLevel,omitempty"`
}
