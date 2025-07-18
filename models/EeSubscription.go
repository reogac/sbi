/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscription struct {
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	Gpsi                       string                             `json:"gpsi,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
	CallbackReference          string                             `json:"callbackReference"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
}
