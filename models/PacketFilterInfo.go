/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
