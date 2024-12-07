/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtAmfEventSubscription struct {
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	NfId                          string                              `json:"nfId"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	Pei                           string                              `json:"pei,omitempty"`
	GroupId                       string                              `json:"groupId,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	Supi                          string                              `json:"supi,omitempty"`
}
