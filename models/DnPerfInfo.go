/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DnPerfInfo struct {
	Snssai     *Snssai  `json:"snssai,omitempty"`
	DnPerf     []DnPerf `json:"dnPerf"`
	Confidence *int     `json:"confidence,omitempty"`
	AppId      string   `json:"appId,omitempty"`
	Dnn        string   `json:"dnn,omitempty"`
}
