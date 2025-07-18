/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtAmfEventSubscription struct {
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	NfId                          string                              `json:"nfId"`
	Pei                           string                              `json:"pei,omitempty"`
	GroupId                       string                              `json:"groupId,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	Supi                          string                              `json:"supi,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
}
