/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:16 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionResourceModifyIndication struct {
	SessionId int16  `json:"sessionId"`
	Transfer  []byte `json:"transfer,omitempty"`
}
