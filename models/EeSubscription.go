/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscription struct {
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
	CallbackReference          string                             `json:"callbackReference"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
	Gpsi                       string                             `json:"gpsi,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
}
