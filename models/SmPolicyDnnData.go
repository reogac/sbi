/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDnnData struct {
	AllowedServices     []string                          `json:"allowedServices,omitempty"`
	Offline             *bool                             `json:"offline,omitempty"`
	LocRoutNotAllowed   *bool                             `json:"locRoutNotAllowed,omitempty"`
	SubscCats           []string                          `json:"subscCats,omitempty"`
	AdcSupport          *bool                             `json:"adcSupport,omitempty"`
	RefUmDataLimitIds   map[string]LimitIdToMonitoringKey `json:"refUmDataLimitIds,omitempty"`
	MpsPriority         *bool                             `json:"mpsPriority,omitempty"`
	McsPriorityLevel    *int                              `json:"mcsPriorityLevel,omitempty"`
	Dnn                 string                            `json:"dnn"`
	GbrDl               string                            `json:"gbrDl,omitempty"`
	SubscSpendingLimits *bool                             `json:"subscSpendingLimits,omitempty"`
	Ipv4Index           *int                              `json:"ipv4Index,omitempty"`
	ChfInfo             *ChargingInformation              `json:"chfInfo,omitempty"`
	BdtRefIds           map[string]string                 `json:"bdtRefIds,omitempty"`
	GbrUl               string                            `json:"gbrUl,omitempty"`
	Ipv6Index           *int                              `json:"ipv6Index,omitempty"`
	Online              *bool                             `json:"online,omitempty"`
	McsPriority         *bool                             `json:"mcsPriority,omitempty"`
	ImsSignallingPrio   *bool                             `json:"imsSignallingPrio,omitempty"`
	MpsPriorityLevel    *int                              `json:"mpsPriorityLevel,omitempty"`
	PraInfos            map[string]PresenceInfo           `json:"praInfos,omitempty"`
}
