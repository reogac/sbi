/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:38 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfUeContextInfo struct {
	AmfSet     string `json:"amfSet"`
	AmfPointer int16  `json:"amfPointer"`
	PlmnId     PlmnId `json:"plmnId"`
	AmfUeId    int64  `json:"amfUeId"`
}
