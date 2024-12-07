/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SubscriptionDataSubscriptions struct {
	SdmSubscription           *SdmSubscription     `json:"sdmSubscription,omitempty"`
	HssSubscriptionInfo       *HssSubscriptionInfo `json:"hssSubscriptionInfo,omitempty"`
	UniqueSubscription        *bool                `json:"uniqueSubscription,omitempty"`
	SupportedFeatures         string               `json:"supportedFeatures,omitempty"`
	CallbackReference         string               `json:"callbackReference"`
	OriginalCallbackReference string               `json:"originalCallbackReference,omitempty"`
	Expiry                    string               `json:"expiry,omitempty"`
	UeId                      string               `json:"ueId,omitempty"`
	MonitoredResourceUris     []string             `json:"monitoredResourceUris"`
	SubscriptionId            string               `json:"subscriptionId,omitempty"`
}
