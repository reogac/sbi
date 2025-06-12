/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SubscriptionDataSubscriptions struct {
	CallbackReference         string               `json:"callbackReference"`
	OriginalCallbackReference string               `json:"originalCallbackReference,omitempty"`
	UniqueSubscription        *bool                `json:"uniqueSubscription,omitempty"`
	SupportedFeatures         string               `json:"supportedFeatures,omitempty"`
	UeId                      string               `json:"ueId,omitempty"`
	MonitoredResourceUris     []string             `json:"monitoredResourceUris"`
	Expiry                    string               `json:"expiry,omitempty"`
	SdmSubscription           *SdmSubscription     `json:"sdmSubscription,omitempty"`
	HssSubscriptionInfo       *HssSubscriptionInfo `json:"hssSubscriptionInfo,omitempty"`
	SubscriptionId            string               `json:"subscriptionId,omitempty"`
}
