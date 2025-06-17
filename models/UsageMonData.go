/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UsageMonData struct {
	ResetTime    string                       `json:"resetTime,omitempty"`
	SuppFeat     string                       `json:"suppFeat,omitempty"`
	ResetIds     []string                     `json:"resetIds,omitempty"`
	LimitId      string                       `json:"limitId"`
	Scopes       map[string]UsageMonDataScope `json:"scopes,omitempty"`
	UmLevel      UsageMonLevel                `json:"umLevel,omitempty"`
	AllowedUsage *UsageThreshold              `json:"allowedUsage,omitempty"`
}
