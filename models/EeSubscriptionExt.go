/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscriptionExt struct {
	Gpsi                       string                             `json:"gpsi,omitempty"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	AmfSubscriptionInfoList    []AmfSubscriptionInfo              `json:"amfSubscriptionInfoList,omitempty"`
	HssSubscriptionInfo        *HssSubscriptionInfo               `json:"hssSubscriptionInfo,omitempty"`
	CallbackReference          string                             `json:"callbackReference"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	SmfSubscriptionInfo        *SmfSubscriptionInfo               `json:"smfSubscriptionInfo,omitempty"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
}
