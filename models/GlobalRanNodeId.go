/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:49:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GlobalRanNodeId struct {
	Nid     string `json:"nid,omitempty"`
	ENbId   string `json:"eNbId,omitempty"`
	PlmnId  PlmnId `json:"plmnId"`
	N3IwfId string `json:"n3IwfId,omitempty"`
	GNbId   *GNbId `json:"gNbId,omitempty"`
	NgeNbId string `json:"ngeNbId,omitempty"`
	WagfId  string `json:"wagfId,omitempty"`
	TngfId  string `json:"tngfId,omitempty"`
}
