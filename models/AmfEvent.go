/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEvent struct {
	Type                   AmfEventType            `json:"type"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
}
