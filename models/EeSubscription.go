/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscription struct {
	CallbackReference          string                             `json:"callbackReference"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
	Gpsi                       string                             `json:"gpsi,omitempty"`
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
}
