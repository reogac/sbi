/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DataChangeNotify struct {
	NotifyItems                   []NotifyItem                    `json:"notifyItems,omitempty"`
	SdmSubscription               *SdmSubscription                `json:"sdmSubscription,omitempty"`
	AdditionalSdmSubscriptions    []SdmSubscription               `json:"additionalSdmSubscriptions,omitempty"`
	SubscriptionDataSubscriptions []SubscriptionDataSubscriptions `json:"subscriptionDataSubscriptions,omitempty"`
	OriginalCallbackReference     []string                        `json:"originalCallbackReference,omitempty"`
	UeId                          string                          `json:"ueId,omitempty"`
}
