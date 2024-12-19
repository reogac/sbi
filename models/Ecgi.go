/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 14:25:56 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Ecgi struct {
	PlmnId      PlmnId `json:"plmnId"`
	EutraCellId string `json:"eutraCellId"`
	Nid         string `json:"nid,omitempty"`
}
