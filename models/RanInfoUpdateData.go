/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:34 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RanInfoUpdateData struct {
	Id              string            `json:"id"`
	PlmnId          PlmnId            `json:"plmnId"`
	SupportedTAList []SupportedTAItem `json:"supportedTAList,omitempty"`
}
