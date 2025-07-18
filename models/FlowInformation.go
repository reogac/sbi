/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type FlowInformation struct {
	FlowDescription    string              `json:"flowDescription,omitempty"`
	EthFlowDescription *EthFlowDescription `json:"ethFlowDescription,omitempty"`
	PackFiltId         string              `json:"packFiltId,omitempty"`
	PacketFilterUsage  *bool               `json:"packetFilterUsage,omitempty"`
	TosTrafficClass    string              `json:"tosTrafficClass,omitempty"`
	Spi                string              `json:"spi,omitempty"`
	FlowLabel          string              `json:"flowLabel,omitempty"`
	FlowDirection      FlowDirection       `json:"flowDirection,omitempty"`
}
