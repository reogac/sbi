/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DispersionCollection struct {
	UeLoc          *UserLocation       `json:"ueLoc,omitempty"`
	Gpsis          []string            `json:"gpsis,omitempty"`
	DisperClass    DispersionClass     `json:"disperClass,omitempty"`
	UsageRank      *int                `json:"usageRank,omitempty"`
	UeRatio        *int                `json:"ueRatio,omitempty"`
	Snssai         *Snssai             `json:"snssai,omitempty"`
	Supis          []string            `json:"supis,omitempty"`
	AppVolumes     []ApplicationVolume `json:"appVolumes,omitempty"`
	DisperAmount   *int                `json:"disperAmount,omitempty"`
	PercentileRank *int                `json:"percentileRank,omitempty"`
	Confidence     *int                `json:"confidence,omitempty"`
}
