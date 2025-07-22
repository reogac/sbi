/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscription struct {
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	Gpsi                       string                             `json:"gpsi,omitempty"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
	CallbackReference          string                             `json:"callbackReference"`
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
}
