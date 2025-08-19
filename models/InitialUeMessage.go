/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type InitialUeMessage struct {
	AuthCtx        *UeAuthCtx        `json:"authCtx,omitempty"`
	NasPdu         []byte            `json:"nasPdu,omitempty"`
	Transfer       bool              `json:"transfer"`
	RrcCause       int16             `json:"rrcCause"`
	Loc            *UserLocation     `json:"loc,omitempty"`
	ContextRequest bool              `json:"contextRequest"`
	Access         AccessType        `json:"access"`
	RanNets        []string          `json:"ranNets,omitempty"`
	RanUeId        RanUeId           `json:"ranUeId"`
	NasSplit       bool              `json:"nasSplit"`
	NfSelection    map[string]string `json:"nfSelection,omitempty"`
	AmfRegion      int16             `json:"amfRegion"`
}
