/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscription struct {
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
	Gpsi                       string                             `json:"gpsi,omitempty"`
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	CallbackReference          string                             `json:"callbackReference"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
}
