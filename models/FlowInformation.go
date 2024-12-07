/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:30 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type FlowInformation struct {
	EthFlowDescription *EthFlowDescription `json:"ethFlowDescription,omitempty"`
	PackFiltId         string              `json:"packFiltId,omitempty"`
	PacketFilterUsage  *bool               `json:"packetFilterUsage,omitempty"`
	TosTrafficClass    string              `json:"tosTrafficClass,omitempty"`
	Spi                string              `json:"spi,omitempty"`
	FlowLabel          string              `json:"flowLabel,omitempty"`
	FlowDirection      FlowDirection       `json:"flowDirection,omitempty"`
	FlowDescription    string              `json:"flowDescription,omitempty"`
}
