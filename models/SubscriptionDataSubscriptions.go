/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SubscriptionDataSubscriptions struct {
	SubscriptionId            string               `json:"subscriptionId,omitempty"`
	UniqueSubscription        *bool                `json:"uniqueSubscription,omitempty"`
	SupportedFeatures         string               `json:"supportedFeatures,omitempty"`
	UeId                      string               `json:"ueId,omitempty"`
	OriginalCallbackReference string               `json:"originalCallbackReference,omitempty"`
	MonitoredResourceUris     []string             `json:"monitoredResourceUris"`
	Expiry                    string               `json:"expiry,omitempty"`
	HssSubscriptionInfo       *HssSubscriptionInfo `json:"hssSubscriptionInfo,omitempty"`
	CallbackReference         string               `json:"callbackReference"`
	SdmSubscription           *SdmSubscription     `json:"sdmSubscription,omitempty"`
}
