/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscription struct {
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	CallbackReference          string                             `json:"callbackReference"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
	Gpsi                       string                             `json:"gpsi,omitempty"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
}
