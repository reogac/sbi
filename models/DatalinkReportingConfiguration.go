/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DatalinkReportingConfiguration struct {
	DddTrafficDes []DddTrafficDescriptor `json:"dddTrafficDes,omitempty"`
	Dnn           string                 `json:"dnn,omitempty"`
	Slice         *Snssai                `json:"slice,omitempty"`
	DddStatusList []string               `json:"dddStatusList,omitempty"`
}
