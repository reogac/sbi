/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:24 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionResourceSetupRequest struct {
	N1Sm      []byte       `json:"n1Sm,omitempty"`
	Ref       string       `json:"ref"`
	Smf       EndpointInfo `json:"smf"`
	Snssai    Snssai       `json:"snssai"`
	SessionId int16        `json:"sessionId"`
	Transfer  []byte       `json:"transfer,omitempty"`
}
