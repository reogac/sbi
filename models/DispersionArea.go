/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DispersionArea struct {
	TaiList  []Tai  `json:"taiList,omitempty"`
	NcgiList []Ncgi `json:"ncgiList,omitempty"`
	EcgiList []Ecgi `json:"ecgiList,omitempty"`
	N3gaInd  *bool  `json:"n3gaInd,omitempty"`
}
