/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:13 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type CircumstanceDescription struct {
	Freq    *float64         `json:"freq,omitempty"`
	Tm      string           `json:"tm,omitempty"`
	LocArea *NetworkAreaInfo `json:"locArea,omitempty"`
	Vol     *int64           `json:"vol,omitempty"`
}
