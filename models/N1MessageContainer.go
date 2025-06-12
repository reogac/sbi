/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1MessageContainer struct {
	N1MessageClass    N1MessageClass  `json:"n1MessageClass"`
	N1MessageContent  RefToBinaryData `json:"n1MessageContent"`
	NfId              string          `json:"nfId,omitempty"`
	ServiceInstanceId string          `json:"serviceInstanceId,omitempty"`
}
