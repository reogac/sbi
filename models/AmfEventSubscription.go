/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEventSubscription struct {
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	Supi                          string        `json:"supi,omitempty"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	NfId                          string        `json:"nfId"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
}
