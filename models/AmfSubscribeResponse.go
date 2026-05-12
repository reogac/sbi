/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:36 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfSubscribeResponse struct {
	Id              string            `json:"id"`
	PlmnId          PlmnId            `json:"plmnId"`
	SupportedTAList []SupportedTAItem `json:"supportedTAList,omitempty"`
}
