/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:22 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionResourceSetupRequest struct {
	Transfer  []byte       `json:"transfer,omitempty"`
	N1Sm      []byte       `json:"n1Sm,omitempty"`
	Ref       string       `json:"ref"`
	Smf       EndpointInfo `json:"smf"`
	Snssai    Snssai       `json:"snssai"`
	SessionId int16        `json:"sessionId"`
}
