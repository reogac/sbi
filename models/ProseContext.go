/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProseContext struct {
	DirectDiscovery UeAuth      `json:"directDiscovery,omitempty"`
	DirectComm      UeAuth      `json:"directComm,omitempty"`
	L2Relay         UeAuth      `json:"l2Relay,omitempty"`
	L3Relay         UeAuth      `json:"l3Relay,omitempty"`
	L2Remote        UeAuth      `json:"l2Remote,omitempty"`
	NrUePc5Ambr     string      `json:"nrUePc5Ambr,omitempty"`
	Pc5QoSPara      *Pc5QoSPara `json:"pc5QoSPara,omitempty"`
}
