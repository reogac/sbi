/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccuUsageReport struct {
	VolUsage             *int64 `json:"volUsage,omitempty"`
	VolUsageDownlink     *int64 `json:"volUsageDownlink,omitempty"`
	NextVolUsageUplink   *int64 `json:"nextVolUsageUplink,omitempty"`
	NextVolUsageDownlink *int64 `json:"nextVolUsageDownlink,omitempty"`
	NextTimeUsage        *int   `json:"nextTimeUsage,omitempty"`
	RefUmIds             string `json:"refUmIds"`
	VolUsageUplink       *int64 `json:"volUsageUplink,omitempty"`
	TimeUsage            *int   `json:"timeUsage,omitempty"`
	NextVolUsage         *int64 `json:"nextVolUsage,omitempty"`
}
