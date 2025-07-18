/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEvent struct {
	Type                   AmfEventType            `json:"type"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
}
