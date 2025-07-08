/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type InitialUeMessage struct {
	AuthCtx        *UeAuthCtx        `json:"authCtx,omitempty"`
	NasPdu         []byte            `json:"nasPdu,omitempty"`
	ContextRequest bool              `json:"contextRequest"`
	Transfer       bool              `json:"transfer"`
	RrcCause       int16             `json:"rrcCause"`
	Loc            *UserLocation     `json:"loc,omitempty"`
	Access         AccessType        `json:"access"`
	NfSelection    map[string]string `json:"nfSelection,omitempty"`
	RanNets        []string          `json:"ranNets,omitempty"`
	RanUeId        RanUeId           `json:"ranUeId"`
	NasSplit       bool              `json:"nasSplit"`
}
