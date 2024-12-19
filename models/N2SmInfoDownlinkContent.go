/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 14:26:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2SmInfoDownlinkContent struct {
	N2SmInfoType N2SmInfoType `json:"n2SmInfoType,omitempty"`
	Snssai       *Snssai      `json:"snssai,omitempty"`
	SessionId    int16        `json:"sessionId"`
	N2SmInfo     []byte       `json:"n2SmInfo,omitempty"`
	NasPdu       []byte       `json:"nasPdu,omitempty"`
}
