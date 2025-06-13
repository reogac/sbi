/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SubscriptionDataSubscriptions struct {
	UeId                      string               `json:"ueId,omitempty"`
	CallbackReference         string               `json:"callbackReference"`
	MonitoredResourceUris     []string             `json:"monitoredResourceUris"`
	Expiry                    string               `json:"expiry,omitempty"`
	SdmSubscription           *SdmSubscription     `json:"sdmSubscription,omitempty"`
	SupportedFeatures         string               `json:"supportedFeatures,omitempty"`
	OriginalCallbackReference string               `json:"originalCallbackReference,omitempty"`
	HssSubscriptionInfo       *HssSubscriptionInfo `json:"hssSubscriptionInfo,omitempty"`
	SubscriptionId            string               `json:"subscriptionId,omitempty"`
	UniqueSubscription        *bool                `json:"uniqueSubscription,omitempty"`
}
