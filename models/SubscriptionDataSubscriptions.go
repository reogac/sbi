/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SubscriptionDataSubscriptions struct {
	CallbackReference         string               `json:"callbackReference"`
	Expiry                    string               `json:"expiry,omitempty"`
	SdmSubscription           *SdmSubscription     `json:"sdmSubscription,omitempty"`
	HssSubscriptionInfo       *HssSubscriptionInfo `json:"hssSubscriptionInfo,omitempty"`
	UniqueSubscription        *bool                `json:"uniqueSubscription,omitempty"`
	SupportedFeatures         string               `json:"supportedFeatures,omitempty"`
	UeId                      string               `json:"ueId,omitempty"`
	OriginalCallbackReference string               `json:"originalCallbackReference,omitempty"`
	MonitoredResourceUris     []string             `json:"monitoredResourceUris"`
	SubscriptionId            string               `json:"subscriptionId,omitempty"`
}
