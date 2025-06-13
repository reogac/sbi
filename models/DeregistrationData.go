/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:23 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DeregistrationData struct {
	DeregReason      DeregistrationReason `json:"deregReason"`
	AccessType       AccessType           `json:"accessType,omitempty"`
	PduSessionId     *int                 `json:"pduSessionId,omitempty"`
	NewSmfInstanceId string               `json:"newSmfInstanceId,omitempty"`
}
