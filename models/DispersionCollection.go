/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DispersionCollection struct {
	UeLoc          *UserLocation       `json:"ueLoc,omitempty"`
	Snssai         *Snssai             `json:"snssai,omitempty"`
	Gpsis          []string            `json:"gpsis,omitempty"`
	AppVolumes     []ApplicationVolume `json:"appVolumes,omitempty"`
	DisperClass    DispersionClass     `json:"disperClass,omitempty"`
	UsageRank      *int                `json:"usageRank,omitempty"`
	Supis          []string            `json:"supis,omitempty"`
	DisperAmount   *int                `json:"disperAmount,omitempty"`
	PercentileRank *int                `json:"percentileRank,omitempty"`
	UeRatio        *int                `json:"ueRatio,omitempty"`
	Confidence     *int                `json:"confidence,omitempty"`
}
