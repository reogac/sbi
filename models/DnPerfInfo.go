/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DnPerfInfo struct {
	Confidence *int     `json:"confidence,omitempty"`
	AppId      string   `json:"appId,omitempty"`
	Dnn        string   `json:"dnn,omitempty"`
	Snssai     *Snssai  `json:"snssai,omitempty"`
	DnPerf     []DnPerf `json:"dnPerf"`
}
