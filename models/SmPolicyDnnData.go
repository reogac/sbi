/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDnnData struct {
	GbrUl               string                            `json:"gbrUl,omitempty"`
	GbrDl               string                            `json:"gbrDl,omitempty"`
	MpsPriorityLevel    *int                              `json:"mpsPriorityLevel,omitempty"`
	PraInfos            map[string]PresenceInfo           `json:"praInfos,omitempty"`
	Dnn                 string                            `json:"dnn"`
	AdcSupport          *bool                             `json:"adcSupport,omitempty"`
	SubscSpendingLimits *bool                             `json:"subscSpendingLimits,omitempty"`
	Ipv4Index           *int                              `json:"ipv4Index,omitempty"`
	McsPriorityLevel    *int                              `json:"mcsPriorityLevel,omitempty"`
	LocRoutNotAllowed   *bool                             `json:"locRoutNotAllowed,omitempty"`
	Ipv6Index           *int                              `json:"ipv6Index,omitempty"`
	ChfInfo             *ChargingInformation              `json:"chfInfo,omitempty"`
	MpsPriority         *bool                             `json:"mpsPriority,omitempty"`
	McsPriority         *bool                             `json:"mcsPriority,omitempty"`
	ImsSignallingPrio   *bool                             `json:"imsSignallingPrio,omitempty"`
	BdtRefIds           map[string]string                 `json:"bdtRefIds,omitempty"`
	AllowedServices     []string                          `json:"allowedServices,omitempty"`
	SubscCats           []string                          `json:"subscCats,omitempty"`
	Offline             *bool                             `json:"offline,omitempty"`
	Online              *bool                             `json:"online,omitempty"`
	RefUmDataLimitIds   map[string]LimitIdToMonitoringKey `json:"refUmDataLimitIds,omitempty"`
}
