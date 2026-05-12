/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:46 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDnnData struct {
	SubscCats           []string                          `json:"subscCats,omitempty"`
	Online              *bool                             `json:"online,omitempty"`
	ChfInfo             *ChargingInformation              `json:"chfInfo,omitempty"`
	RefUmDataLimitIds   map[string]LimitIdToMonitoringKey `json:"refUmDataLimitIds,omitempty"`
	PraInfos            map[string]PresenceInfo           `json:"praInfos,omitempty"`
	LocRoutNotAllowed   *bool                             `json:"locRoutNotAllowed,omitempty"`
	GbrUl               string                            `json:"gbrUl,omitempty"`
	SubscSpendingLimits *bool                             `json:"subscSpendingLimits,omitempty"`
	Ipv4Index           *int                              `json:"ipv4Index,omitempty"`
	Ipv6Index           *int                              `json:"ipv6Index,omitempty"`
	ImsSignallingPrio   *bool                             `json:"imsSignallingPrio,omitempty"`
	MpsPriorityLevel    *int                              `json:"mpsPriorityLevel,omitempty"`
	AllowedServices     []string                          `json:"allowedServices,omitempty"`
	GbrDl               string                            `json:"gbrDl,omitempty"`
	AdcSupport          *bool                             `json:"adcSupport,omitempty"`
	Offline             *bool                             `json:"offline,omitempty"`
	MpsPriority         *bool                             `json:"mpsPriority,omitempty"`
	McsPriority         *bool                             `json:"mcsPriority,omitempty"`
	McsPriorityLevel    *int                              `json:"mcsPriorityLevel,omitempty"`
	BdtRefIds           map[string]string                 `json:"bdtRefIds,omitempty"`
	Dnn                 string                            `json:"dnn"`
}
