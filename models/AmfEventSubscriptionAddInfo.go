/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEventSubscriptionAddInfo struct {
	AoiStateList      map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	BindingInfo       []string                            `json:"bindingInfo,omitempty"`
	SubscribingNfType NFType                              `json:"subscribingNfType,omitempty"`
	EventSyncInd      *bool                               `json:"eventSyncInd,omitempty"`
	NfConsumerInfo    []string                            `json:"nfConsumerInfo,omitempty"`
}
