/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DnPerfInfo struct {
	AppId      string   `json:"appId,omitempty"`
	Dnn        string   `json:"dnn,omitempty"`
	Snssai     *Snssai  `json:"snssai,omitempty"`
	DnPerf     []DnPerf `json:"dnPerf"`
	Confidence *int     `json:"confidence,omitempty"`
}
