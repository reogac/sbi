/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:21 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionResourceNotification struct {
	Release   bool   `json:"release"`
	Transfer  []byte `json:"transfer,omitempty"`
	SessionId int16  `json:"sessionId"`
}
