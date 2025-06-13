/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:18 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AlternativeQosProfile struct {
	PacketErrRate     string `json:"packetErrRate,omitempty"`
	Index             int    `json:"index"`
	GuaFbrDl          string `json:"guaFbrDl,omitempty"`
	GuaFbrUl          string `json:"guaFbrUl,omitempty"`
	PacketDelayBudget *int   `json:"packetDelayBudget,omitempty"`
}
