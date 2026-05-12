/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:34 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfUeContextInfo struct {
	AmfUeId    int64  `json:"amfUeId"`
	AmfSet     string `json:"amfSet"`
	AmfPointer int16  `json:"amfPointer"`
	PlmnId     PlmnId `json:"plmnId"`
}
