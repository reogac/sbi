/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Feb  7 10:27:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RanInfoUpdateData struct {
	PlmnId          PlmnId            `json:"plmnId"`
	SupportedTAList []SupportedTAItem `json:"supportedTAList,omitempty"`
	Id              string            `json:"id"`
}
