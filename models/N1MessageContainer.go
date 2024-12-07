/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1MessageContainer struct {
	NfId              string          `json:"nfId,omitempty"`
	ServiceInstanceId string          `json:"serviceInstanceId,omitempty"`
	N1MessageClass    N1MessageClass  `json:"n1MessageClass"`
	N1MessageContent  RefToBinaryData `json:"n1MessageContent"`
}
