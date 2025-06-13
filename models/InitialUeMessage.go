/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:34 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type InitialUeMessage struct {
	RanUeId        RanUeId           `json:"ranUeId"`
	NasSplit       bool              `json:"nasSplit"`
	NfSelection    map[string]string `json:"nfSelection,omitempty"`
	NasPdu         []byte            `json:"nasPdu,omitempty"`
	Access         AccessType        `json:"access"`
	Transfer       bool              `json:"transfer"`
	RrcCause       int16             `json:"rrcCause"`
	Loc            *UserLocation     `json:"loc,omitempty"`
	RanNets        []string          `json:"ranNets,omitempty"`
	AuthCtx        *UeAuthCtx        `json:"authCtx,omitempty"`
	ContextRequest bool              `json:"contextRequest"`
}
