/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:01 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type CellGlobalId struct {
	CellId string `json:"cellId"`
	PlmnId PlmnId `json:"plmnId"`
	Lac    string `json:"lac"`
}
