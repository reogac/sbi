/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:47 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionResourceNotification struct {
	Release   bool   `json:"release"`
	Transfer  []byte `json:"transfer,omitempty"`
	SessionId int16  `json:"sessionId"`
}
