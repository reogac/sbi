/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:30 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PacketFilterInfo struct {
	PackFiltId      string        `json:"packFiltId,omitempty"`
	PackFiltCont    string        `json:"packFiltCont,omitempty"`
	TosTrafficClass string        `json:"tosTrafficClass,omitempty"`
	Spi             string        `json:"spi,omitempty"`
	FlowLabel       string        `json:"flowLabel,omitempty"`
	FlowDirection   FlowDirection `json:"flowDirection,omitempty"`
}
