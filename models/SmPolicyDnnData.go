/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:38 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDnnData struct {
	ChfInfo             *ChargingInformation              `json:"chfInfo,omitempty"`
	PraInfos            map[string]PresenceInfo           `json:"praInfos,omitempty"`
	GbrUl               string                            `json:"gbrUl,omitempty"`
	Ipv4Index           *int                              `json:"ipv4Index,omitempty"`
	GbrDl               string                            `json:"gbrDl,omitempty"`
	AdcSupport          *bool                             `json:"adcSupport,omitempty"`
	SubscSpendingLimits *bool                             `json:"subscSpendingLimits,omitempty"`
	ImsSignallingPrio   *bool                             `json:"imsSignallingPrio,omitempty"`
	MpsPriorityLevel    *int                              `json:"mpsPriorityLevel,omitempty"`
	McsPriorityLevel    *int                              `json:"mcsPriorityLevel,omitempty"`
	Dnn                 string                            `json:"dnn"`
	SubscCats           []string                          `json:"subscCats,omitempty"`
	MpsPriority         *bool                             `json:"mpsPriority,omitempty"`
	BdtRefIds           map[string]string                 `json:"bdtRefIds,omitempty"`
	LocRoutNotAllowed   *bool                             `json:"locRoutNotAllowed,omitempty"`
	AllowedServices     []string                          `json:"allowedServices,omitempty"`
	Offline             *bool                             `json:"offline,omitempty"`
	RefUmDataLimitIds   map[string]LimitIdToMonitoringKey `json:"refUmDataLimitIds,omitempty"`
	McsPriority         *bool                             `json:"mcsPriority,omitempty"`
	Ipv6Index           *int                              `json:"ipv6Index,omitempty"`
	Online              *bool                             `json:"online,omitempty"`
}
