/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSession struct {
	PlmnId        PlmnId  `json:"plmnId"`
	SingleNssai   *Snssai `json:"singleNssai,omitempty"`
	Dnn           string  `json:"dnn"`
	SmfInstanceId string  `json:"smfInstanceId"`
}
