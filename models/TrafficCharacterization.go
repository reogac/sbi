/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TrafficCharacterization struct {
	UlVol         *int64                 `json:"ulVol,omitempty"`
	UlVolVariance *float64               `json:"ulVolVariance,omitempty"`
	DlVol         *int64                 `json:"dlVol,omitempty"`
	DlVolVariance *float64               `json:"dlVolVariance,omitempty"`
	Dnn           string                 `json:"dnn,omitempty"`
	Snssai        *Snssai                `json:"snssai,omitempty"`
	AppId         string                 `json:"appId,omitempty"`
	FDescs        []IpEthFlowDescription `json:"fDescs,omitempty"`
}
