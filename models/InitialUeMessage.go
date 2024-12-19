/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:49:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type InitialUeMessage struct {
	ContextRequest bool              `json:"contextRequest"`
	RrcCause       int16             `json:"rrcCause"`
	Loc            *UserLocation     `json:"loc,omitempty"`
	RanNets        []string          `json:"ranNets,omitempty"`
	RanUeId        RanUeId           `json:"ranUeId"`
	NfSelection    map[string]string `json:"nfSelection,omitempty"`
	AuthCtx        *UeAuthCtx        `json:"authCtx,omitempty"`
	Transfer       bool              `json:"transfer"`
	Access         AccessType        `json:"access"`
	NasSplit       bool              `json:"nasSplit"`
	NasPdu         []byte            `json:"nasPdu,omitempty"`
}
