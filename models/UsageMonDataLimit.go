/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UsageMonDataLimit struct {
	UmLevel     UsageMonLevel                `json:"umLevel,omitempty"`
	StartDate   string                       `json:"startDate,omitempty"`
	EndDate     string                       `json:"endDate,omitempty"`
	UsageLimit  *UsageThreshold              `json:"usageLimit,omitempty"`
	ResetPeriod *TimePeriod                  `json:"resetPeriod,omitempty"`
	LimitId     string                       `json:"limitId"`
	Scopes      map[string]UsageMonDataScope `json:"scopes,omitempty"`
}
