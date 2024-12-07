/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:21 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2SmInfoDownlinkTransport struct {
	Transfers []N2SmInfoDownlinkContent `json:"transfers,omitempty"`
	NasPdu    []byte                    `json:"nasPdu,omitempty"`
}
