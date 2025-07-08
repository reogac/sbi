/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtAmfEventSubscription struct {
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	NfId                          string                              `json:"nfId"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	Pei                           string                              `json:"pei,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	GroupId                       string                              `json:"groupId,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	Supi                          string                              `json:"supi,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
}
