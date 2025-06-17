/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextCreatedData struct {
	PcfReselectedInd     *bool             `json:"pcfReselectedInd,omitempty"`
	AnalyticsNotUsedList []string          `json:"analyticsNotUsedList,omitempty"`
	UeContext            UeContext         `json:"ueContext"`
	TargetToSourceData   N2InfoContent     `json:"targetToSourceData"`
	PduSessionList       []N2SmInformation `json:"pduSessionList"`
	FailedSessionList    []N2SmInformation `json:"failedSessionList,omitempty"`
	SupportedFeatures    string            `json:"supportedFeatures,omitempty"`
}
