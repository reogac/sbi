/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:13 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtAmfEventSubscription struct {
	GroupId                       string                              `json:"groupId,omitempty"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	Pei                           string                              `json:"pei,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	NfId                          string                              `json:"nfId"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	Supi                          string                              `json:"supi,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
}
