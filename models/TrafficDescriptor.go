/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TrafficDescriptor struct {
	DddTrafficDescriptorList []DddTrafficDescriptor `json:"dddTrafficDescriptorList,omitempty"`
	Dnn                      string                 `json:"dnn,omitempty"`
	SNssai                   *Snssai                `json:"sNssai,omitempty"`
}
