/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscriptionExt struct {
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	SmfSubscriptionInfo        *SmfSubscriptionInfo               `json:"smfSubscriptionInfo,omitempty"`
	HssSubscriptionInfo        *HssSubscriptionInfo               `json:"hssSubscriptionInfo,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	CallbackReference          string                             `json:"callbackReference"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	AmfSubscriptionInfoList    []AmfSubscriptionInfo              `json:"amfSubscriptionInfoList,omitempty"`
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	Gpsi                       string                             `json:"gpsi,omitempty"`
}
