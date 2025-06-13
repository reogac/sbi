/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:50 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UePolicyTransferFailureNotification struct {
	Ptis  []int                    `json:"ptis"`
	Cause N1N2MessageTransferCause `json:"cause"`
}
