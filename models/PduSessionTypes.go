/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionTypes struct {
	DefaultSessionType  PduSessionType `json:"defaultSessionType,omitempty"`
	AllowedSessionTypes []string       `json:"allowedSessionTypes,omitempty"`
}
