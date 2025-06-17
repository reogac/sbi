/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GlobalRanNodeId struct {
	GNbId   *GNbId `json:"gNbId,omitempty"`
	NgeNbId string `json:"ngeNbId,omitempty"`
	WagfId  string `json:"wagfId,omitempty"`
	TngfId  string `json:"tngfId,omitempty"`
	Nid     string `json:"nid,omitempty"`
	ENbId   string `json:"eNbId,omitempty"`
	PlmnId  PlmnId `json:"plmnId"`
	N3IwfId string `json:"n3IwfId,omitempty"`
}
