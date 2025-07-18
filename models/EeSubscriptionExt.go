/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscriptionExt struct {
	SmfSubscriptionInfo        *SmfSubscriptionInfo               `json:"smfSubscriptionInfo,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	CallbackReference          string                             `json:"callbackReference"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	Gpsi                       string                             `json:"gpsi,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	HssSubscriptionInfo        *HssSubscriptionInfo               `json:"hssSubscriptionInfo,omitempty"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	AmfSubscriptionInfoList    []AmfSubscriptionInfo              `json:"amfSubscriptionInfoList,omitempty"`
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
}
