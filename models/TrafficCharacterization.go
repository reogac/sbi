/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TrafficCharacterization struct {
	AppId         string                 `json:"appId,omitempty"`
	FDescs        []IpEthFlowDescription `json:"fDescs,omitempty"`
	UlVol         *int64                 `json:"ulVol,omitempty"`
	UlVolVariance *float64               `json:"ulVolVariance,omitempty"`
	DlVol         *int64                 `json:"dlVol,omitempty"`
	DlVolVariance *float64               `json:"dlVolVariance,omitempty"`
	Dnn           string                 `json:"dnn,omitempty"`
	Snssai        *Snssai                `json:"snssai,omitempty"`
}
