/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:13 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type V2xContext struct {
	Pc5QoSPara         *Pc5QoSPara `json:"pc5QoSPara,omitempty"`
	NrV2xServicesAuth  *NrV2xAuth  `json:"nrV2xServicesAuth,omitempty"`
	LteV2xServicesAuth *LteV2xAuth `json:"lteV2xServicesAuth,omitempty"`
	NrUeSidelinkAmbr   string      `json:"nrUeSidelinkAmbr,omitempty"`
	LteUeSidelinkAmbr  string      `json:"lteUeSidelinkAmbr,omitempty"`
}
