/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionResourceModifyRequest struct {
	Transfer  []byte `json:"transfer,omitempty"`
	N1Sm      []byte `json:"n1Sm,omitempty"`
	SessionId int16  `json:"sessionId"`
}
