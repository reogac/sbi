/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1MessageContainer struct {
	N1MessageClass    N1MessageClass  `json:"n1MessageClass"`
	N1MessageContent  RefToBinaryData `json:"n1MessageContent"`
	NfId              string          `json:"nfId,omitempty"`
	ServiceInstanceId string          `json:"serviceInstanceId,omitempty"`
}
