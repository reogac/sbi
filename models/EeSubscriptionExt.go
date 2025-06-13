/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscriptionExt struct {
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
	Gpsi                       string                             `json:"gpsi,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	AmfSubscriptionInfoList    []AmfSubscriptionInfo              `json:"amfSubscriptionInfoList,omitempty"`
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	SmfSubscriptionInfo        *SmfSubscriptionInfo               `json:"smfSubscriptionInfo,omitempty"`
	HssSubscriptionInfo        *HssSubscriptionInfo               `json:"hssSubscriptionInfo,omitempty"`
	CallbackReference          string                             `json:"callbackReference"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
}
