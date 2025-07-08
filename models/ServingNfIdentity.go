/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:44 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServingNfIdentity struct {
	ServNfInstId string       `json:"servNfInstId,omitempty"`
	Guami        *Guami       `json:"guami,omitempty"`
	AnGwAddr     *AnGwAddress `json:"anGwAddr,omitempty"`
	SgsnAddr     *SgsnAddress `json:"sgsnAddr,omitempty"`
}
