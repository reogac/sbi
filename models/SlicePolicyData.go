/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SlicePolicyData struct {
	MbrUl       string   `json:"mbrUl,omitempty"`
	MbrDl       string   `json:"mbrDl,omitempty"`
	RemainMbrUl string   `json:"remainMbrUl,omitempty"`
	RemainMbrDl string   `json:"remainMbrDl,omitempty"`
	SuppFeat    string   `json:"suppFeat,omitempty"`
	ResetIds    []string `json:"resetIds,omitempty"`
}
