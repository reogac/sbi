/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UsageMonDataLimit struct {
	UsageLimit  *UsageThreshold              `json:"usageLimit,omitempty"`
	ResetPeriod *TimePeriod                  `json:"resetPeriod,omitempty"`
	LimitId     string                       `json:"limitId"`
	Scopes      map[string]UsageMonDataScope `json:"scopes,omitempty"`
	UmLevel     UsageMonLevel                `json:"umLevel,omitempty"`
	StartDate   string                       `json:"startDate,omitempty"`
	EndDate     string                       `json:"endDate,omitempty"`
}
