/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:44 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type InitialUeMessage struct {
	Access         AccessType        `json:"access"`
	NasSplit       bool              `json:"nasSplit"`
	NasPdu         []byte            `json:"nasPdu,omitempty"`
	Transfer       bool              `json:"transfer"`
	Loc            *UserLocation     `json:"loc,omitempty"`
	RanNets        []string          `json:"ranNets,omitempty"`
	RanUeId        RanUeId           `json:"ranUeId"`
	NfSelection    map[string]string `json:"nfSelection,omitempty"`
	AuthCtx        *UeAuthCtx        `json:"authCtx,omitempty"`
	ContextRequest bool              `json:"contextRequest"`
	RrcCause       int16             `json:"rrcCause"`
}
