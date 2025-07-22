/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfUeContextInfo struct {
	AmfSet     string `json:"amfSet"`
	AmfPointer int16  `json:"amfPointer"`
	PlmnId     PlmnId `json:"plmnId"`
	AmfUeId    int64  `json:"amfUeId"`
}
