/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:24 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionResourceModifyRequest struct {
	SessionId int16  `json:"sessionId"`
	Transfer  []byte `json:"transfer,omitempty"`
	N1Sm      []byte `json:"n1Sm,omitempty"`
}
