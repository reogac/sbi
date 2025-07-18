/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextCreatedData struct {
	UeContext            UeContext         `json:"ueContext"`
	TargetToSourceData   N2InfoContent     `json:"targetToSourceData"`
	PduSessionList       []N2SmInformation `json:"pduSessionList"`
	FailedSessionList    []N2SmInformation `json:"failedSessionList,omitempty"`
	SupportedFeatures    string            `json:"supportedFeatures,omitempty"`
	PcfReselectedInd     *bool             `json:"pcfReselectedInd,omitempty"`
	AnalyticsNotUsedList []string          `json:"analyticsNotUsedList,omitempty"`
}
