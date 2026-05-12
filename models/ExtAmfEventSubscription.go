/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtAmfEventSubscription struct {
	Pei                           string                              `json:"pei,omitempty"`
	AnyUE                         *bool                               `json:"anyUE,omitempty"`
	NfId                          string                              `json:"nfId"`
	GroupId                       string                              `json:"groupId,omitempty"`
	IncludeGpsiList               []string                            `json:"includeGpsiList,omitempty"`
	EventSyncInd                  *bool                               `json:"eventSyncInd,omitempty"`
	NotifyCorrelationId           string                              `json:"notifyCorrelationId"`
	Supi                          string                              `json:"supi,omitempty"`
	ExcludeGpsiList               []string                            `json:"excludeGpsiList,omitempty"`
	NfConsumerInfo                []string                            `json:"nfConsumerInfo,omitempty"`
	ExcludeSupiList               []string                            `json:"excludeSupiList,omitempty"`
	SourceNfType                  NFType                              `json:"sourceNfType,omitempty"`
	Gpsi                          string                              `json:"gpsi,omitempty"`
	EventNotifyUri                string                              `json:"eventNotifyUri"`
	SubscribingNfType             NFType                              `json:"subscribingNfType,omitempty"`
	BindingInfo                   []string                            `json:"bindingInfo,omitempty"`
	SubsChangeNotifyUri           string                              `json:"subsChangeNotifyUri,omitempty"`
	SubsChangeNotifyCorrelationId string                              `json:"subsChangeNotifyCorrelationId,omitempty"`
	IncludeSupiList               []string                            `json:"includeSupiList,omitempty"`
	EventList                     []AmfEvent                          `json:"eventList"`
	Options                       *AmfEventMode                       `json:"options,omitempty"`
	AoiStateList                  map[string]AreaOfInterestEventState `json:"aoiStateList,omitempty"`
}
