/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionResourceNotification struct {
	Transfer  []byte `json:"transfer,omitempty"`
	SessionId int16  `json:"sessionId"`
	Release   bool   `json:"release"`
}
