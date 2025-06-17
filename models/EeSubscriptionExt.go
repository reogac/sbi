/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscriptionExt struct {
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	SmfSubscriptionInfo        *SmfSubscriptionInfo               `json:"smfSubscriptionInfo,omitempty"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	Gpsi                       string                             `json:"gpsi,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	AmfSubscriptionInfoList    []AmfSubscriptionInfo              `json:"amfSubscriptionInfoList,omitempty"`
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	CallbackReference          string                             `json:"callbackReference"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	HssSubscriptionInfo        *HssSubscriptionInfo               `json:"hssSubscriptionInfo,omitempty"`
}
