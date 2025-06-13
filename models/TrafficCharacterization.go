/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TrafficCharacterization struct {
	UlVolVariance *float64               `json:"ulVolVariance,omitempty"`
	DlVol         *int64                 `json:"dlVol,omitempty"`
	DlVolVariance *float64               `json:"dlVolVariance,omitempty"`
	Dnn           string                 `json:"dnn,omitempty"`
	Snssai        *Snssai                `json:"snssai,omitempty"`
	AppId         string                 `json:"appId,omitempty"`
	FDescs        []IpEthFlowDescription `json:"fDescs,omitempty"`
	UlVol         *int64                 `json:"ulVol,omitempty"`
}
