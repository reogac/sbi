/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PrevSubInfo struct {
	ProducerSetId  string                         `json:"producerSetId,omitempty"`
	SubscriptionId string                         `json:"subscriptionId"`
	NfAnaEvents    []string                       `json:"nfAnaEvents,omitempty"`
	UeAnaEvents    []UeAnalyticsContextDescriptor `json:"ueAnaEvents,omitempty"`
	ProducerId     string                         `json:"producerId,omitempty"`
}
