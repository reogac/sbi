/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:24 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfUeContextInfo struct {
	AmfSet     string `json:"amfSet"`
	AmfPointer int16  `json:"amfPointer"`
	PlmnId     PlmnId `json:"plmnId"`
	AmfUeId    int64  `json:"amfUeId"`
}
