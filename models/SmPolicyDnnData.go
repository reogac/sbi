/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:34 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDnnData struct {
	Online              *bool                             `json:"online,omitempty"`
	RefUmDataLimitIds   map[string]LimitIdToMonitoringKey `json:"refUmDataLimitIds,omitempty"`
	ImsSignallingPrio   *bool                             `json:"imsSignallingPrio,omitempty"`
	BdtRefIds           map[string]string                 `json:"bdtRefIds,omitempty"`
	SubscCats           []string                          `json:"subscCats,omitempty"`
	GbrDl               string                            `json:"gbrDl,omitempty"`
	AdcSupport          *bool                             `json:"adcSupport,omitempty"`
	Ipv6Index           *int                              `json:"ipv6Index,omitempty"`
	McsPriorityLevel    *int                              `json:"mcsPriorityLevel,omitempty"`
	PraInfos            map[string]PresenceInfo           `json:"praInfos,omitempty"`
	GbrUl               string                            `json:"gbrUl,omitempty"`
	Ipv4Index           *int                              `json:"ipv4Index,omitempty"`
	MpsPriority         *bool                             `json:"mpsPriority,omitempty"`
	McsPriority         *bool                             `json:"mcsPriority,omitempty"`
	LocRoutNotAllowed   *bool                             `json:"locRoutNotAllowed,omitempty"`
	Dnn                 string                            `json:"dnn"`
	AllowedServices     []string                          `json:"allowedServices,omitempty"`
	SubscSpendingLimits *bool                             `json:"subscSpendingLimits,omitempty"`
	Offline             *bool                             `json:"offline,omitempty"`
	ChfInfo             *ChargingInformation              `json:"chfInfo,omitempty"`
	MpsPriorityLevel    *int                              `json:"mpsPriorityLevel,omitempty"`
}
