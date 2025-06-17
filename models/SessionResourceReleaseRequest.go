/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionResourceReleaseRequest struct {
	N1Sm      []byte `json:"n1Sm,omitempty"`
	SessionId int16  `json:"sessionId"`
	Transfer  []byte `json:"transfer,omitempty"`
}
