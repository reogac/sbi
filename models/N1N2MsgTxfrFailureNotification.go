/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N1N2MsgTxfrFailureNotification struct {
	Cause          N1N2MessageTransferCause `json:"cause"`
	N1n2MsgDataUri string                   `json:"n1n2MsgDataUri"`
}
