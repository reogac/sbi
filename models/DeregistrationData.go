/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DeregistrationData struct {
	NewSmfInstanceId string               `json:"newSmfInstanceId,omitempty"`
	DeregReason      DeregistrationReason `json:"deregReason"`
	AccessType       AccessType           `json:"accessType,omitempty"`
	PduSessionId     *int                 `json:"pduSessionId,omitempty"`
}
