/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscriptionExt struct {
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
	Gpsi                       string                             `json:"gpsi,omitempty"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	SmfSubscriptionInfo        *SmfSubscriptionInfo               `json:"smfSubscriptionInfo,omitempty"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	HssSubscriptionInfo        *HssSubscriptionInfo               `json:"hssSubscriptionInfo,omitempty"`
	CallbackReference          string                             `json:"callbackReference"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	AmfSubscriptionInfoList    []AmfSubscriptionInfo              `json:"amfSubscriptionInfoList,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
}
