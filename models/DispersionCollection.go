/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DispersionCollection struct {
	UsageRank      *int                `json:"usageRank,omitempty"`
	UeRatio        *int                `json:"ueRatio,omitempty"`
	Confidence     *int                `json:"confidence,omitempty"`
	Snssai         *Snssai             `json:"snssai,omitempty"`
	Supis          []string            `json:"supis,omitempty"`
	Gpsis          []string            `json:"gpsis,omitempty"`
	PercentileRank *int                `json:"percentileRank,omitempty"`
	UeLoc          *UserLocation       `json:"ueLoc,omitempty"`
	AppVolumes     []ApplicationVolume `json:"appVolumes,omitempty"`
	DisperAmount   *int                `json:"disperAmount,omitempty"`
	DisperClass    DispersionClass     `json:"disperClass,omitempty"`
}
