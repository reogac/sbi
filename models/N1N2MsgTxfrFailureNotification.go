/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1N2MsgTxfrFailureNotification struct {
	N1n2MsgDataUri string                   `json:"n1n2MsgDataUri"`
	Cause          N1N2MessageTransferCause `json:"cause"`
}
