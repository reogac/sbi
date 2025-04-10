/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DnnInfo struct {
	DefaultDnnIndicator *bool    `json:"defaultDnnIndicator,omitempty"`
	LboRoamingAllowed   *bool    `json:"lboRoamingAllowed,omitempty"`
	IwkEpsInd           *bool    `json:"iwkEpsInd,omitempty"`
	DnnBarred           *bool    `json:"dnnBarred,omitempty"`
	InvokeNefInd        *bool    `json:"invokeNefInd,omitempty"`
	SmfList             []string `json:"smfList,omitempty"`
	SameSmfInd          *bool    `json:"sameSmfInd,omitempty"`
	Dnn                 string   `json:"dnn"`
}
