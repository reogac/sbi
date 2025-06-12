/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:18 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type InitialUeMessage struct {
	RrcCause       int16             `json:"rrcCause"`
	Loc            *UserLocation     `json:"loc,omitempty"`
	RanNets        []string          `json:"ranNets,omitempty"`
	NasSplit       bool              `json:"nasSplit"`
	AuthCtx        *UeAuthCtx        `json:"authCtx,omitempty"`
	Transfer       bool              `json:"transfer"`
	Access         AccessType        `json:"access"`
	RanUeId        RanUeId           `json:"ranUeId"`
	NfSelection    map[string]string `json:"nfSelection,omitempty"`
	NasPdu         []byte            `json:"nasPdu,omitempty"`
	ContextRequest bool              `json:"contextRequest"`
}
