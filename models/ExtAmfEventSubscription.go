/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtAmfEventSubscription struct {
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	NfId                          string                              `json:"nfId"`
	Pei                           string                              `json:"pei,omitempty"`
	GroupId                       string                              `json:"groupId,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	Supi                          string                              `json:"supi,omitempty"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
}
