/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtAmfEventSubscription struct {
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	NfId                          string                              `json:"nfId"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	Supi                          string                              `json:"supi,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	Pei                           string                              `json:"pei,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	GroupId                       string                              `json:"groupId,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
}
