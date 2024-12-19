/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:49:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AreaScope struct {
	TacList         []string           `json:"tacList,omitempty"`
	TacInfoPerPlmn  map[string]TacInfo `json:"tacInfoPerPlmn,omitempty"`
	EutraCellIdList []string           `json:"eutraCellIdList,omitempty"`
	NrCellIdList    []string           `json:"nrCellIdList,omitempty"`
}
