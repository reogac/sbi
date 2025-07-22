/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscriptionExt struct {
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	Gpsi                       string                             `json:"gpsi,omitempty"`
	CallbackReference          string                             `json:"callbackReference"`
	HssSubscriptionInfo        *HssSubscriptionInfo               `json:"hssSubscriptionInfo,omitempty"`
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	SmfSubscriptionInfo        *SmfSubscriptionInfo               `json:"smfSubscriptionInfo,omitempty"`
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
	AmfSubscriptionInfoList    []AmfSubscriptionInfo              `json:"amfSubscriptionInfoList,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
}
