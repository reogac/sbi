/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Feb 18 15:01:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2SmInfoUplinkContent struct {
	SessionId    int16        `json:"sessionId"`
	N2SmInfo     []byte       `json:"n2SmInfo,omitempty"`
	N2SmInfoType N2SmInfoType `json:"n2SmInfoType,omitempty"`
}
