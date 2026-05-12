/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:28 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type InitialUeMessage struct {
	AuthCtx        *UeAuthCtx        `json:"authCtx,omitempty"`
	NasPdu         []byte            `json:"nasPdu,omitempty"`
	Transfer       bool              `json:"transfer"`
	RrcCause       int16             `json:"rrcCause"`
	Loc            *UserLocation     `json:"loc,omitempty"`
	RanUeId        RanUeId           `json:"ranUeId"`
	ContextRequest bool              `json:"contextRequest"`
	Access         AccessType        `json:"access"`
	RanNets        []string          `json:"ranNets,omitempty"`
	NasSplit       bool              `json:"nasSplit"`
	NfSelection    map[string]string `json:"nfSelection,omitempty"`
	AmfRegion      int16             `json:"amfRegion"`
}
