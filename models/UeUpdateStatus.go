/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeUpdateStatus string

// Define constant values for UeUpdateStatus
const (
	UEUPDATESTATUS_NOT_SENT              UeUpdateStatus = "NOT_SENT"
	UEUPDATESTATUS_SENT_NO_ACK_REQUIRED  UeUpdateStatus = "SENT_NO_ACK_REQUIRED"
	UEUPDATESTATUS_WAITING_FOR_ACK       UeUpdateStatus = "WAITING_FOR_ACK"
	UEUPDATESTATUS_ACK_RECEIVED          UeUpdateStatus = "ACK_RECEIVED"
	UEUPDATESTATUS_NEGATIVE_ACK_RECEIVED UeUpdateStatus = "NEGATIVE_ACK_RECEIVED"
)
