/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DispersionArea struct {
	EcgiList []Ecgi `json:"ecgiList,omitempty"`
	N3gaInd  *bool  `json:"n3gaInd,omitempty"`
	TaiList  []Tai  `json:"taiList,omitempty"`
	NcgiList []Ncgi `json:"ncgiList,omitempty"`
}
