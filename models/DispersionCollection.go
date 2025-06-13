/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DispersionCollection struct {
	Snssai         *Snssai             `json:"snssai,omitempty"`
	Gpsis          []string            `json:"gpsis,omitempty"`
	AppVolumes     []ApplicationVolume `json:"appVolumes,omitempty"`
	UsageRank      *int                `json:"usageRank,omitempty"`
	PercentileRank *int                `json:"percentileRank,omitempty"`
	Confidence     *int                `json:"confidence,omitempty"`
	UeLoc          *UserLocation       `json:"ueLoc,omitempty"`
	Supis          []string            `json:"supis,omitempty"`
	DisperAmount   *int                `json:"disperAmount,omitempty"`
	DisperClass    DispersionClass     `json:"disperClass,omitempty"`
	UeRatio        *int                `json:"ueRatio,omitempty"`
}
