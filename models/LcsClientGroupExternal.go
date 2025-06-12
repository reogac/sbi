/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type LcsClientGroupExternal struct {
	PrivacyCheckRelatedAction PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty"`
	ValidTimePeriod           *ValidTimePeriod          `json:"validTimePeriod,omitempty"`
	LcsClientGroupId          string                    `json:"lcsClientGroupId,omitempty"`
	AllowedGeographicArea     []GeographicArea          `json:"allowedGeographicArea,omitempty"`
}
