/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RanInfoUpdateData struct {
	Id              string            `json:"id"`
	PlmnId          PlmnId            `json:"plmnId"`
	SupportedTAList []SupportedTAItem `json:"supportedTAList,omitempty"`
}
