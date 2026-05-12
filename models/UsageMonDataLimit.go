/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:46 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UsageMonDataLimit struct {
	LimitId     string                       `json:"limitId"`
	Scopes      map[string]UsageMonDataScope `json:"scopes,omitempty"`
	UmLevel     UsageMonLevel                `json:"umLevel,omitempty"`
	StartDate   string                       `json:"startDate,omitempty"`
	EndDate     string                       `json:"endDate,omitempty"`
	UsageLimit  *UsageThreshold              `json:"usageLimit,omitempty"`
	ResetPeriod *TimePeriod                  `json:"resetPeriod,omitempty"`
}
