/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RegistrationLocationInfo struct {
	VgmlcAddress   *VgmlcAddress `json:"vgmlcAddress,omitempty"`
	AccessTypeList []string      `json:"accessTypeList"`
	AmfInstanceId  string        `json:"amfInstanceId"`
	Guami          *Guami        `json:"guami,omitempty"`
	PlmnId         *PlmnId       `json:"plmnId,omitempty"`
}
