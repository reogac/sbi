/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEvent struct {
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	Type                   AmfEventType            `json:"type"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
}
