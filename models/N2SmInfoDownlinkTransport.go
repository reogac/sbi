/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:32 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2SmInfoDownlinkTransport struct {
	Transfers []N2SmInfoDownlinkContent `json:"transfers,omitempty"`
	NasPdu    []byte                    `json:"nasPdu,omitempty"`
}
