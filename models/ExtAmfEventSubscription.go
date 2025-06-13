/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtAmfEventSubscription struct {
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	Supi                          string                              `json:"supi,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	NfId                          string                              `json:"nfId"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	Pei                           string                              `json:"pei,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	GroupId                       string                              `json:"groupId,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
}
