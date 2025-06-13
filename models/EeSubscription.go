/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscription struct {
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
	CallbackReference          string                             `json:"callbackReference"`
	Gpsi                       string                             `json:"gpsi,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
}
