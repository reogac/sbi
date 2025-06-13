/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscription struct {
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
	CallbackReference          string                             `json:"callbackReference"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	Gpsi                       string                             `json:"gpsi,omitempty"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
}
