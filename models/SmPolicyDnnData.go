/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:50 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDnnData struct {
	LocRoutNotAllowed   *bool                             `json:"locRoutNotAllowed,omitempty"`
	AdcSupport          *bool                             `json:"adcSupport,omitempty"`
	RefUmDataLimitIds   map[string]LimitIdToMonitoringKey `json:"refUmDataLimitIds,omitempty"`
	MpsPriority         *bool                             `json:"mpsPriority,omitempty"`
	MpsPriorityLevel    *int                              `json:"mpsPriorityLevel,omitempty"`
	McsPriorityLevel    *int                              `json:"mcsPriorityLevel,omitempty"`
	PraInfos            map[string]PresenceInfo           `json:"praInfos,omitempty"`
	BdtRefIds           map[string]string                 `json:"bdtRefIds,omitempty"`
	AllowedServices     []string                          `json:"allowedServices,omitempty"`
	Dnn                 string                            `json:"dnn"`
	SubscCats           []string                          `json:"subscCats,omitempty"`
	GbrDl               string                            `json:"gbrDl,omitempty"`
	ChfInfo             *ChargingInformation              `json:"chfInfo,omitempty"`
	ImsSignallingPrio   *bool                             `json:"imsSignallingPrio,omitempty"`
	GbrUl               string                            `json:"gbrUl,omitempty"`
	SubscSpendingLimits *bool                             `json:"subscSpendingLimits,omitempty"`
	Ipv4Index           *int                              `json:"ipv4Index,omitempty"`
	Ipv6Index           *int                              `json:"ipv6Index,omitempty"`
	Offline             *bool                             `json:"offline,omitempty"`
	Online              *bool                             `json:"online,omitempty"`
	McsPriority         *bool                             `json:"mcsPriority,omitempty"`
}
