/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:21 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionResourceSetupRequest struct {
	Smf       EndpointInfo `json:"smf"`
	Snssai    Snssai       `json:"snssai"`
	SessionId int16        `json:"sessionId"`
	Transfer  []byte       `json:"transfer,omitempty"`
	N1Sm      []byte       `json:"n1Sm,omitempty"`
	Ref       string       `json:"ref"`
}