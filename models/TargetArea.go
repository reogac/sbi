/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TargetArea struct {
	TaList       []Tai      `json:"taList,omitempty"`
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
	AnyTa        *bool      `json:"anyTa,omitempty"`
}
