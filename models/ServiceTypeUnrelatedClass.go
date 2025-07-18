/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceTypeUnrelatedClass struct {
	AllowedGeographicArea     []GeographicArea          `json:"allowedGeographicArea,omitempty"`
	PrivacyCheckRelatedAction PrivacyCheckRelatedAction `json:"privacyCheckRelatedAction,omitempty"`
	CodeWordInd               CodeWordInd               `json:"codeWordInd,omitempty"`
	ValidTimePeriod           *ValidTimePeriod          `json:"validTimePeriod,omitempty"`
	CodeWordList              []string                  `json:"codeWordList,omitempty"`
	ServiceType               int                       `json:"serviceType"`
}
