/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtAmfEventSubscription struct {
	Gpsi                          string                              `json:"gpsi,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	Supi                          string                              `json:"supi,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	GroupId                       string                              `json:"groupId,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	Pei                           string                              `json:"pei,omitempty"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	NfId                          string                              `json:"nfId"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
}
