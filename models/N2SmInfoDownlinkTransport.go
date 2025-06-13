/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type N2SmInfoDownlinkTransport struct {
	Transfers []N2SmInfoDownlinkContent `json:"transfers,omitempty"`
	NasPdu    []byte                    `json:"nasPdu,omitempty"`
}
