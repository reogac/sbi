/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:52 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PagingMessage struct {
	PagingOrigin *bool `json:"pagingOrigin,omitempty"`
	TaList       []Tai `json:"taList,omitempty"`
	Tmsi         int32 `json:"tmsi"`
	AmfSet       int16 `json:"amfSet"`
	AmfPointer   int16 `json:"amfPointer"`
}
