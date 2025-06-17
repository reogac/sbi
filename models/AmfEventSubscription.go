/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEventSubscription struct {
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	Supi                          string        `json:"supi,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	GroupId                       string        `json:"groupId,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	NfId                          string        `json:"nfId"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
}
