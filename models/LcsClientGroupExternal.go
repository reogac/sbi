/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type LcsClientGroupExternal struct {
	LcsClientGroupId          string                    `json:"lcsClientGroupId,omitempty"`
	AllowedGeographicArea     []GeographicArea          `json:"allowedGeographicArea,omitempty"`
	PrivacyCheckRelatedAction PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty"`
	ValidTimePeriod           *ValidTimePeriod          `json:"validTimePeriod,omitempty"`
}
