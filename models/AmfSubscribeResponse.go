/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfSubscribeResponse struct {
	Id              string            `json:"id"`
	PlmnId          PlmnId            `json:"plmnId"`
	SupportedTAList []SupportedTAItem `json:"supportedTAList,omitempty"`
}
