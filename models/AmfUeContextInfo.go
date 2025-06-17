/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:50 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfUeContextInfo struct {
	AmfPointer int16  `json:"amfPointer"`
	PlmnId     PlmnId `json:"plmnId"`
	AmfUeId    int64  `json:"amfUeId"`
	AmfSet     string `json:"amfSet"`
}
