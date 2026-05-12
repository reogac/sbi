/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:42 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PacketFilterInfo struct {
	PackFiltCont    string        `json:"packFiltCont,omitempty"`
	TosTrafficClass string        `json:"tosTrafficClass,omitempty"`
	Spi             string        `json:"spi,omitempty"`
	FlowLabel       string        `json:"flowLabel,omitempty"`
	FlowDirection   FlowDirection `json:"flowDirection,omitempty"`
	PackFiltId      string        `json:"packFiltId,omitempty"`
}
