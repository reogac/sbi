/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEvent struct {
	MaxResponseTime        *int                    `json:"maxResponseTime,omitempty"`
	MinInterval            *int                    `json:"minInterval,omitempty"`
	NextReport             string                  `json:"nextReport,omitempty"`
	ImmediateFlag          *bool                   `json:"immediateFlag,omitempty"`
	ReachabilityFilter     ReachabilityFilter      `json:"reachabilityFilter,omitempty"`
	UdmDetectInd           *bool                   `json:"udmDetectInd,omitempty"`
	MaxReports             *int                    `json:"maxReports,omitempty"`
	AreaList               []AmfEventArea          `json:"areaList,omitempty"`
	TrafficDescriptorList  []TrafficDescriptor     `json:"trafficDescriptorList,omitempty"`
	UeInAreaFilter         *UeInAreaFilter         `json:"ueInAreaFilter,omitempty"`
	NextPeriodicReportTime string                  `json:"nextPeriodicReportTime,omitempty"`
	Type                   AmfEventType            `json:"type"`
	ReportUeReachable      *bool                   `json:"reportUeReachable,omitempty"`
	PresenceInfoList       map[string]PresenceInfo `json:"presenceInfoList,omitempty"`
	TargetArea             *TargetArea             `json:"targetArea,omitempty"`
	DispersionArea         *DispersionArea         `json:"dispersionArea,omitempty"`
	LocationFilterList     []string                `json:"locationFilterList,omitempty"`
	RefId                  *int                    `json:"refId,omitempty"`
	SnssaiFilter           []ExtSnssai             `json:"snssaiFilter,omitempty"`
	IdleStatusInd          *bool                   `json:"idleStatusInd,omitempty"`
}
