/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DnnRouteSelectionDescriptor struct {
	Dnn          string   `json:"dnn"`
	SscModes     []string `json:"sscModes,omitempty"`
	PduSessTypes []string `json:"pduSessTypes,omitempty"`
	AtsssInfo    *bool    `json:"atsssInfo,omitempty"`
}
