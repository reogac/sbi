/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SubscriptionDataSubscriptions struct {
	UeId                      string               `json:"ueId,omitempty"`
	CallbackReference         string               `json:"callbackReference"`
	OriginalCallbackReference string               `json:"originalCallbackReference,omitempty"`
	Expiry                    string               `json:"expiry,omitempty"`
	HssSubscriptionInfo       *HssSubscriptionInfo `json:"hssSubscriptionInfo,omitempty"`
	SupportedFeatures         string               `json:"supportedFeatures,omitempty"`
	MonitoredResourceUris     []string             `json:"monitoredResourceUris"`
	SdmSubscription           *SdmSubscription     `json:"sdmSubscription,omitempty"`
	SubscriptionId            string               `json:"subscriptionId,omitempty"`
	UniqueSubscription        *bool                `json:"uniqueSubscription,omitempty"`
}
