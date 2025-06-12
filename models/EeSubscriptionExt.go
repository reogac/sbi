/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscriptionExt struct {
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
	CallbackReference          string                             `json:"callbackReference"`
	Gpsi                       string                             `json:"gpsi,omitempty"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	HssSubscriptionInfo        *HssSubscriptionInfo               `json:"hssSubscriptionInfo,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	AmfSubscriptionInfoList    []AmfSubscriptionInfo              `json:"amfSubscriptionInfoList,omitempty"`
	SmfSubscriptionInfo        *SmfSubscriptionInfo               `json:"smfSubscriptionInfo,omitempty"`
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
}
