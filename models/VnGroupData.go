/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type VnGroupData struct {
	PduSessionTypes *PduSessionTypes `json:"pduSessionTypes,omitempty"`
	Dnn             string           `json:"dnn,omitempty"`
	SingleNssai     *Snssai          `json:"singleNssai,omitempty"`
	AppDescriptors  []AppDescriptor  `json:"appDescriptors,omitempty"`
}
