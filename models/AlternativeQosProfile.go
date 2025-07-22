/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:22 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AlternativeQosProfile struct {
	Index             int    `json:"index"`
	GuaFbrDl          string `json:"guaFbrDl,omitempty"`
	GuaFbrUl          string `json:"guaFbrUl,omitempty"`
	PacketDelayBudget *int   `json:"packetDelayBudget,omitempty"`
	PacketErrRate     string `json:"packetErrRate,omitempty"`
}
