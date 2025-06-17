/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DataChangeNotify struct {
	OriginalCallbackReference     []string                        `json:"originalCallbackReference,omitempty"`
	UeId                          string                          `json:"ueId,omitempty"`
	NotifyItems                   []NotifyItem                    `json:"notifyItems,omitempty"`
	SdmSubscription               *SdmSubscription                `json:"sdmSubscription,omitempty"`
	AdditionalSdmSubscriptions    []SdmSubscription               `json:"additionalSdmSubscriptions,omitempty"`
	SubscriptionDataSubscriptions []SubscriptionDataSubscriptions `json:"subscriptionDataSubscriptions,omitempty"`
}
