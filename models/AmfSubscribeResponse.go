/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Feb  6 14:02:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfSubscribeResponse struct {
	Id              string            `json:"id"`
	PlmnId          PlmnId            `json:"plmnId"`
	SupportedTAList []SupportedTAItem `json:"supportedTAList,omitempty"`
}
