/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextCreatedData struct {
	FailedSessionList    []N2SmInformation `json:"failedSessionList,omitempty"`
	SupportedFeatures    string            `json:"supportedFeatures,omitempty"`
	PcfReselectedInd     *bool             `json:"pcfReselectedInd,omitempty"`
	AnalyticsNotUsedList []string          `json:"analyticsNotUsedList,omitempty"`
	UeContext            UeContext         `json:"ueContext"`
	TargetToSourceData   N2InfoContent     `json:"targetToSourceData"`
	PduSessionList       []N2SmInformation `json:"pduSessionList"`
}
