/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:46 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UsageMonData struct {
	UmLevel      UsageMonLevel                `json:"umLevel,omitempty"`
	AllowedUsage *UsageThreshold              `json:"allowedUsage,omitempty"`
	ResetTime    string                       `json:"resetTime,omitempty"`
	SuppFeat     string                       `json:"suppFeat,omitempty"`
	ResetIds     []string                     `json:"resetIds,omitempty"`
	LimitId      string                       `json:"limitId"`
	Scopes       map[string]UsageMonDataScope `json:"scopes,omitempty"`
}
