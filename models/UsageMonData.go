/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:34 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UsageMonData struct {
	ResetIds     []string                     `json:"resetIds,omitempty"`
	LimitId      string                       `json:"limitId"`
	Scopes       map[string]UsageMonDataScope `json:"scopes,omitempty"`
	UmLevel      UsageMonLevel                `json:"umLevel,omitempty"`
	AllowedUsage *UsageThreshold              `json:"allowedUsage,omitempty"`
	ResetTime    string                       `json:"resetTime,omitempty"`
	SuppFeat     string                       `json:"suppFeat,omitempty"`
}
