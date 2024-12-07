/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DnPerf struct {
	Dnai             string           `json:"dnai,omitempty"`
	PerfData         PerfData         `json:"perfData"`
	SpatialValidCon  *NetworkAreaInfo `json:"spatialValidCon,omitempty"`
	TemporalValidCon *TimeWindow      `json:"temporalValidCon,omitempty"`
	AppServerInsAddr *AddrFqdn        `json:"appServerInsAddr,omitempty"`
	UpfInfo          *UpfInformation  `json:"upfInfo,omitempty"`
}
