/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
