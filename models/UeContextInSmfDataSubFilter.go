/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextInSmfDataSubFilter struct {
	DnnList      []string `json:"dnnList,omitempty"`
	SnssaiList   []Snssai `json:"snssaiList,omitempty"`
	EmergencyInd *bool    `json:"emergencyInd,omitempty"`
}
