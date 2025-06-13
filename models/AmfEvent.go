/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEvent struct {
	NextReport             string                  `json:"nextReport,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	Type                   AmfEventType            `json:"type"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
}
