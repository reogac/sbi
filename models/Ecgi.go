/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:44:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Ecgi struct {
	Nid         string `json:"nid,omitempty"`
	PlmnId      PlmnId `json:"plmnId"`
	EutraCellId string `json:"eutraCellId"`
}
