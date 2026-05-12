/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:44 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UePolicyTransferFailureNotification struct {
	Ptis  []int                    `json:"ptis"`
	Cause N1N2MessageTransferCause `json:"cause"`
}
