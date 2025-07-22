/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SubscriptionDataSubscriptions struct {
	MonitoredResourceUris     []string             `json:"monitoredResourceUris"`
	HssSubscriptionInfo       *HssSubscriptionInfo `json:"hssSubscriptionInfo,omitempty"`
	SupportedFeatures         string               `json:"supportedFeatures,omitempty"`
	UeId                      string               `json:"ueId,omitempty"`
	OriginalCallbackReference string               `json:"originalCallbackReference,omitempty"`
	SdmSubscription           *SdmSubscription     `json:"sdmSubscription,omitempty"`
	SubscriptionId            string               `json:"subscriptionId,omitempty"`
	UniqueSubscription        *bool                `json:"uniqueSubscription,omitempty"`
	CallbackReference         string               `json:"callbackReference"`
	Expiry                    string               `json:"expiry,omitempty"`
}
