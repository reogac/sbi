/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PrevSubInfo struct {
	UeAnaEvents    []UeAnalyticsContextDescriptor `json:"ueAnaEvents,omitempty"`
	ProducerId     string                         `json:"producerId,omitempty"`
	ProducerSetId  string                         `json:"producerSetId,omitempty"`
	SubscriptionId string                         `json:"subscriptionId"`
	NfAnaEvents    []string                       `json:"nfAnaEvents,omitempty"`
}
