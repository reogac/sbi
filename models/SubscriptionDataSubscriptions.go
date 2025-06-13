/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SubscriptionDataSubscriptions struct {
	CallbackReference         string               `json:"callbackReference"`
	OriginalCallbackReference string               `json:"originalCallbackReference,omitempty"`
	SdmSubscription           *SdmSubscription     `json:"sdmSubscription,omitempty"`
	HssSubscriptionInfo       *HssSubscriptionInfo `json:"hssSubscriptionInfo,omitempty"`
	UeId                      string               `json:"ueId,omitempty"`
	MonitoredResourceUris     []string             `json:"monitoredResourceUris"`
	Expiry                    string               `json:"expiry,omitempty"`
	SubscriptionId            string               `json:"subscriptionId,omitempty"`
	UniqueSubscription        *bool                `json:"uniqueSubscription,omitempty"`
	SupportedFeatures         string               `json:"supportedFeatures,omitempty"`
}
