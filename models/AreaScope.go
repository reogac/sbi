/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AreaScope struct {
	EutraCellIdList []string           `json:"eutraCellIdList,omitempty"`
	NrCellIdList    []string           `json:"nrCellIdList,omitempty"`
	TacList         []string           `json:"tacList,omitempty"`
	TacInfoPerPlmn  map[string]TacInfo `json:"tacInfoPerPlmn,omitempty"`
}
