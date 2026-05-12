/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEventSubscription struct {
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	NfId                          string        `json:"nfId"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	GroupId                       string        `json:"groupId,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	Supi                          string        `json:"supi,omitempty"`
}
