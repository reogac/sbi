/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:48:22 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UpdateAmfInfo struct {
	AmfSet     string `json:"amfSet"`
	AmfPointer int16  `json:"amfPointer"`
	PlmnId     PlmnId `json:"plmnId"`
	AmfUeId    int64  `json:"amfUeId"`
}
