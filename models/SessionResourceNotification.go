/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:31 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionResourceNotification struct {
	SessionId int16  `json:"sessionId"`
	Release   bool   `json:"release"`
	Transfer  []byte `json:"transfer,omitempty"`
}
