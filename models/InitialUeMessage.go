/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:16 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type InitialUeMessage struct {
	NasPdu         []byte            `json:"nasPdu,omitempty"`
	ContextRequest bool              `json:"contextRequest"`
	RrcCause       int16             `json:"rrcCause"`
	RanUeId        RanUeId           `json:"ranUeId"`
	AuthCtx        *UeAuthCtx        `json:"authCtx,omitempty"`
	Transfer       bool              `json:"transfer"`
	Loc            *UserLocation     `json:"loc,omitempty"`
	Access         AccessType        `json:"access"`
	RanNets        []string          `json:"ranNets,omitempty"`
	NasSplit       bool              `json:"nasSplit"`
	NfSelection    map[string]string `json:"nfSelection,omitempty"`
}
