/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEvent struct {
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	Type                   AmfEventType            `json:"type"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
}
