/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEventSubscriptionAddInfo struct {
	BindingInfo       []string                            `json:"bindingInfo,omitempty"`
	SubscribingNfType NFType                              `json:"subscribingNfType,omitempty"`
	EventSyncInd      *bool                               `json:"eventSyncInd,omitempty"`
	NfConsumerInfo    []string                            `json:"nfConsumerInfo,omitempty"`
	AoiStateList      map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
}
