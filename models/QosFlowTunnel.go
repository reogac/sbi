/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type QosFlowTunnel struct {
	TunnelInfo TunnelInfo `json:"tunnelInfo"`
	QfiList    []int      `json:"qfiList"`
}
