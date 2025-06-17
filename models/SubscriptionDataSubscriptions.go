/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SubscriptionDataSubscriptions struct {
	UeId                      string               `json:"ueId,omitempty"`
	CallbackReference         string               `json:"callbackReference"`
	OriginalCallbackReference string               `json:"originalCallbackReference,omitempty"`
	MonitoredResourceUris     []string             `json:"monitoredResourceUris"`
	SdmSubscription           *SdmSubscription     `json:"sdmSubscription,omitempty"`
	Expiry                    string               `json:"expiry,omitempty"`
	HssSubscriptionInfo       *HssSubscriptionInfo `json:"hssSubscriptionInfo,omitempty"`
	SubscriptionId            string               `json:"subscriptionId,omitempty"`
	UniqueSubscription        *bool                `json:"uniqueSubscription,omitempty"`
	SupportedFeatures         string               `json:"supportedFeatures,omitempty"`
}
