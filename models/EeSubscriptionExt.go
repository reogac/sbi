/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EeSubscriptionExt struct {
	SmfSubscriptionInfo        *SmfSubscriptionInfo               `json:"smfSubscriptionInfo,omitempty"`
	HssSubscriptionInfo        *HssSubscriptionInfo               `json:"hssSubscriptionInfo,omitempty"`
	Gpsi                       string                             `json:"gpsi,omitempty"`
	ScefDiamRealm              string                             `json:"scefDiamRealm,omitempty"`
	SubscriptionId             string                             `json:"subscriptionId,omitempty"`
	IncludeGpsiList            []string                           `json:"includeGpsiList,omitempty"`
	MonitoringConfigurations   map[string]MonitoringConfiguration `json:"monitoringConfigurations"`
	ScefDiamHost               string                             `json:"scefDiamHost,omitempty"`
	SecondCallbackRef          string                             `json:"secondCallbackRef,omitempty"`
	ExcludeGpsiList            []string                           `json:"excludeGpsiList,omitempty"`
	UdrRestartInd              *bool                              `json:"udrRestartInd,omitempty"`
	ContextInfo                *ContextInfo                       `json:"contextInfo,omitempty"`
	NotifyCorrelationId        string                             `json:"notifyCorrelationId,omitempty"`
	DataRestorationCallbackUri string                             `json:"dataRestorationCallbackUri,omitempty"`
	SupportedFeatures          string                             `json:"supportedFeatures,omitempty"`
	EpcAppliedInd              *bool                              `json:"epcAppliedInd,omitempty"`
	CallbackReference          string                             `json:"callbackReference"`
	ReportingOptions           *ReportingOptions                  `json:"reportingOptions,omitempty"`
	AmfSubscriptionInfoList    []AmfSubscriptionInfo              `json:"amfSubscriptionInfoList,omitempty"`
}
