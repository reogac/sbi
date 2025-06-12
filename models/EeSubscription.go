/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscription struct {
	Gpsi                       string                             `json:"gpsi,omitempty"`
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	CallbackReference          string                             `json:"callbackReference"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
}
