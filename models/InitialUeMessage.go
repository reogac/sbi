/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:23 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type InitialUeMessage struct {
	RrcCause       int16             `json:"rrcCause"`
	Loc            *UserLocation     `json:"loc,omitempty"`
	Access         AccessType        `json:"access"`
	RanNets        []string          `json:"ranNets,omitempty"`
	RanUeId        RanUeId           `json:"ranUeId"`
	AuthCtx        *UeAuthCtx        `json:"authCtx,omitempty"`
	NasPdu         []byte            `json:"nasPdu,omitempty"`
	Transfer       bool              `json:"transfer"`
	NasSplit       bool              `json:"nasSplit"`
	ContextRequest bool              `json:"contextRequest"`
	NfSelection    map[string]string `json:"nfSelection,omitempty"`
}
