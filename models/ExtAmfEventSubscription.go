/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtAmfEventSubscription struct {
	GroupId                       string                              `json:"groupId,omitempty"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	Pei                           string                              `json:"pei,omitempty"`
	Supi                          string                              `json:"supi,omitempty"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	NfId                          string                              `json:"nfId"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
}
