/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEventSubscription struct {
	Supi                          string        `json:"supi,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	NfId                          string        `json:"nfId"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
}
