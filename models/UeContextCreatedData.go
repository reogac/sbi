/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextCreatedData struct {
	AnalyticsNotUsedList []string          `json:"analyticsNotUsedList,omitempty"`
	UeContext            UeContext         `json:"ueContext"`
	TargetToSourceData   N2InfoContent     `json:"targetToSourceData"`
	PduSessionList       []N2SmInformation `json:"pduSessionList"`
	FailedSessionList    []N2SmInformation `json:"failedSessionList,omitempty"`
	SupportedFeatures    string            `json:"supportedFeatures,omitempty"`
	PcfReselectedInd     *bool             `json:"pcfReselectedInd,omitempty"`
}
