/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscriptionExt struct {
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	HssSubscriptionInfo        *HssSubscriptionInfo               `json:"hssSubscriptionInfo,omitempty"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
	CallbackReference          string                             `json:"callbackReference"`
	SmfSubscriptionInfo        *SmfSubscriptionInfo               `json:"smfSubscriptionInfo,omitempty"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	Gpsi                       string                             `json:"gpsi,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	AmfSubscriptionInfoList    []AmfSubscriptionInfo              `json:"amfSubscriptionInfoList,omitempty"`
}
