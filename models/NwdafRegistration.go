/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:23 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NwdafRegistration struct {
	RegistrationTime  string       `json:"registrationTime,omitempty"`
	ContextInfo       *ContextInfo `json:"contextInfo,omitempty"`
	SupportedFeatures string       `json:"supportedFeatures,omitempty"`
	ResetIds          []string     `json:"resetIds,omitempty"`
	NwdafInstanceId   string       `json:"nwdafInstanceId"`
	AnalyticsIds      []string     `json:"analyticsIds"`
	NwdafSetId        string       `json:"nwdafSetId,omitempty"`
}
