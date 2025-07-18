/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RegistrationLocationInfo struct {
	AmfInstanceId  string        `json:"amfInstanceId"`
	Guami          *Guami        `json:"guami,omitempty"`
	PlmnId         *PlmnId       `json:"plmnId,omitempty"`
	VgmlcAddress   *VgmlcAddress `json:"vgmlcAddress,omitempty"`
	AccessTypeList []string      `json:"accessTypeList"`
}
