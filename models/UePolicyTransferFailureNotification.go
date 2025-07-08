/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UePolicyTransferFailureNotification struct {
	Cause N1N2MessageTransferCause `json:"cause"`
	Ptis  []int                    `json:"ptis"`
}
