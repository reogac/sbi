/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type LcsClientGroupExternal struct {
	PrivacyCheckRelatedAction PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty"`
	ValidTimePeriod           *ValidTimePeriod          `json:"validTimePeriod,omitempty"`
	LcsClientGroupId          string                    `json:"lcsClientGroupId,omitempty"`
	AllowedGeographicArea     []GeographicArea          `json:"allowedGeographicArea,omitempty"`
}
