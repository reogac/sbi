/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProseContext struct {
	Pc5QoSPara      *Pc5QoSPara `json:"pc5QoSPara,omitempty"`
	DirectDiscovery UeAuth      `json:"directDiscovery,omitempty"`
	DirectComm      UeAuth      `json:"directComm,omitempty"`
	L2Relay         UeAuth      `json:"l2Relay,omitempty"`
	L3Relay         UeAuth      `json:"l3Relay,omitempty"`
	L2Remote        UeAuth      `json:"l2Remote,omitempty"`
	NrUePc5Ambr     string      `json:"nrUePc5Ambr,omitempty"`
}
