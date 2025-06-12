/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEvent struct {
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	Type                   AmfEventType            `json:"type"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
}
