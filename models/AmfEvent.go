/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEvent struct {
	MinInterval            *int                    `json:"minInterval,omitempty"`
	Type                   AmfEventType            `json:"type"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
}
