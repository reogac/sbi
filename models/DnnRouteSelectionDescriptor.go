/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:46 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DnnRouteSelectionDescriptor struct {
	AtsssInfo    *bool    `json:"atsssInfo,omitempty"`
	Dnn          string   `json:"dnn"`
	SscModes     []string `json:"sscModes,omitempty"`
	PduSessTypes []string `json:"pduSessTypes,omitempty"`
}
