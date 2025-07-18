/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccuUsageReport struct {
	NextVolUsage         *int64 `json:"nextVolUsage,omitempty"`
	NextVolUsageUplink   *int64 `json:"nextVolUsageUplink,omitempty"`
	VolUsageUplink       *int64 `json:"volUsageUplink,omitempty"`
	VolUsageDownlink     *int64 `json:"volUsageDownlink,omitempty"`
	TimeUsage            *int   `json:"timeUsage,omitempty"`
	NextVolUsageDownlink *int64 `json:"nextVolUsageDownlink,omitempty"`
	NextTimeUsage        *int   `json:"nextTimeUsage,omitempty"`
	RefUmIds             string `json:"refUmIds"`
	VolUsage             *int64 `json:"volUsage,omitempty"`
}
