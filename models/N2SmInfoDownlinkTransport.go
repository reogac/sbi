/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2SmInfoDownlinkTransport struct {
	Transfers []N2SmInfoDownlinkContent `json:"transfers,omitempty"`
	NasPdu    []byte                    `json:"nasPdu,omitempty"`
}
