/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:35 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2SmInfoDownlinkContent struct {
	Snssai       *Snssai      `json:"snssai,omitempty"`
	SessionId    int16        `json:"sessionId"`
	N2SmInfo     []byte       `json:"n2SmInfo,omitempty"`
	NasPdu       []byte       `json:"nasPdu,omitempty"`
	N2SmInfoType N2SmInfoType `json:"n2SmInfoType,omitempty"`
}
