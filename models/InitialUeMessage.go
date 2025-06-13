/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:13 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type InitialUeMessage struct {
	Transfer       bool              `json:"transfer"`
	RrcCause       int16             `json:"rrcCause"`
	Access         AccessType        `json:"access"`
	NfSelection    map[string]string `json:"nfSelection,omitempty"`
	ContextRequest bool              `json:"contextRequest"`
	NasPdu         []byte            `json:"nasPdu,omitempty"`
	Loc            *UserLocation     `json:"loc,omitempty"`
	RanNets        []string          `json:"ranNets,omitempty"`
	RanUeId        RanUeId           `json:"ranUeId"`
	NasSplit       bool              `json:"nasSplit"`
	AuthCtx        *UeAuthCtx        `json:"authCtx,omitempty"`
}
