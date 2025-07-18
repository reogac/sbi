/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:50 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Ecgi struct {
	PlmnId      PlmnId `json:"plmnId"`
	EutraCellId string `json:"eutraCellId"`
	Nid         string `json:"nid,omitempty"`
}
