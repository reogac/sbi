/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEventSubscription struct {
	Gpsi                          string        `json:"gpsi,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	Supi                          string        `json:"supi,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	NfId                          string        `json:"nfId"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
}
