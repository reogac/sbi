/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:20 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionResourceNotification struct {
	SessionId int16  `json:"sessionId"`
	Release   bool   `json:"release"`
	Transfer  []byte `json:"transfer,omitempty"`
}
