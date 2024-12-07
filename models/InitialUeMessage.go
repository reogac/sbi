/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:17 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type InitialUeMessage struct {
	Transfer       bool              `json:"transfer"`
	Loc            *UserLocation     `json:"loc,omitempty"`
	RanUeId        RanUeId           `json:"ranUeId"`
	NasPdu         []byte            `json:"nasPdu,omitempty"`
	ContextRequest bool              `json:"contextRequest"`
	Access         AccessType        `json:"access"`
	RanNets        []string          `json:"ranNets,omitempty"`
	NasSplit       bool              `json:"nasSplit"`
	NfSelection    map[string]string `json:"nfSelection,omitempty"`
	AuthCtx        *UeAuthCtx        `json:"authCtx,omitempty"`
	RrcCause       int16             `json:"rrcCause"`
}
