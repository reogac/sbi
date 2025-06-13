/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:13 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEvent struct {
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	Type                   AmfEventType            `json:"type"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
}
