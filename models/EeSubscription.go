/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscription struct {
	CallbackReference          string                             `json:"callbackReference"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	Gpsi                       string                             `json:"gpsi,omitempty"`
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
}
