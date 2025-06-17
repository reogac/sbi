/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtAmfEventSubscription struct {
	Pei                           string                              `json:"pei,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	NfId                          string                              `json:"nfId"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	GroupId                       string                              `json:"groupId,omitempty"`
	Supi                          string                              `json:"supi,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
}
