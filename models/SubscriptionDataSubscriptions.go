/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SubscriptionDataSubscriptions struct {
	CallbackReference         string               `json:"callbackReference"`
	OriginalCallbackReference string               `json:"originalCallbackReference,omitempty"`
	MonitoredResourceUris     []string             `json:"monitoredResourceUris"`
	SdmSubscription           *SdmSubscription     `json:"sdmSubscription,omitempty"`
	HssSubscriptionInfo       *HssSubscriptionInfo `json:"hssSubscriptionInfo,omitempty"`
	SupportedFeatures         string               `json:"supportedFeatures,omitempty"`
	UeId                      string               `json:"ueId,omitempty"`
	Expiry                    string               `json:"expiry,omitempty"`
	SubscriptionId            string               `json:"subscriptionId,omitempty"`
	UniqueSubscription        *bool                `json:"uniqueSubscription,omitempty"`
}
