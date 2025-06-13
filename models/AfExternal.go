/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:21 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AfExternal struct {
	ValidTimePeriod           *ValidTimePeriod          `json:"validTimePeriod,omitempty"`
	AfId                      string                    `json:"afId,omitempty"`
	AllowedGeographicArea     []GeographicArea          `json:"allowedGeographicArea,omitempty"`
	PrivacyCheckRelatedAction PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty"`
}
