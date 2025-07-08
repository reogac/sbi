/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:47 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Ncgi struct {
	Nid      string `json:"nid,omitempty"`
	PlmnId   PlmnId `json:"plmnId"`
	NrCellId string `json:"nrCellId"`
}
