/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscription struct {
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	Gpsi                       string                             `json:"gpsi,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
	CallbackReference          string                             `json:"callbackReference"`
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
}
