/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PagingMessage struct {
	Tmsi         int32 `json:"tmsi"`
	AmfSet       int16 `json:"amfSet"`
	AmfPointer   int16 `json:"amfPointer"`
	PagingOrigin *bool `json:"pagingOrigin,omitempty"`
	TaList       []Tai `json:"taList,omitempty"`
}
