/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:37 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type LcsClientGroupExternal struct {
	LcsClientGroupId          string                    `json:"lcsClientGroupId,omitempty"`
	AllowedGeographicArea     []GeographicArea          `json:"allowedGeographicArea,omitempty"`
	PrivacyCheckRelatedAction PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty"`
	ValidTimePeriod           *ValidTimePeriod          `json:"validTimePeriod,omitempty"`
}
