/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type InitialUeMessage struct {
	RanNets        []string          `json:"ranNets,omitempty"`
	NasSplit       bool              `json:"nasSplit"`
	NfSelection    map[string]string `json:"nfSelection,omitempty"`
	AuthCtx        *UeAuthCtx        `json:"authCtx,omitempty"`
	NasPdu         []byte            `json:"nasPdu,omitempty"`
	Transfer       bool              `json:"transfer"`
	Access         AccessType        `json:"access"`
	RanUeId        RanUeId           `json:"ranUeId"`
	ContextRequest bool              `json:"contextRequest"`
	RrcCause       int16             `json:"rrcCause"`
	Loc            *UserLocation     `json:"loc,omitempty"`
}
