/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DnPerf struct {
	PerfData         PerfData         `json:"perfData"`
	SpatialValidCon  *NetworkAreaInfo `json:"spatialValidCon,omitempty"`
	TemporalValidCon *TimeWindow      `json:"temporalValidCon,omitempty"`
	AppServerInsAddr *AddrFqdn        `json:"appServerInsAddr,omitempty"`
	UpfInfo          *UpfInformation  `json:"upfInfo,omitempty"`
	Dnai             string           `json:"dnai,omitempty"`
}
