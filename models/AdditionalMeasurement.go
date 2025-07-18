/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AdditionalMeasurement struct {
	UnexpWakes    []string                  `json:"unexpWakes,omitempty"`
	DdosAttack    *AddressList              `json:"ddosAttack,omitempty"`
	WrgDest       *AddressList              `json:"wrgDest,omitempty"`
	Circums       []CircumstanceDescription `json:"circums,omitempty"`
	UnexpLoc      *NetworkAreaInfo          `json:"unexpLoc,omitempty"`
	UnexpFlowTeps []IpEthFlowDescription    `json:"unexpFlowTeps,omitempty"`
}
