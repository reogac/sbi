/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEventSubscription struct {
	EventNotifyUri                string        `json:"eventNotifyUri"`
	NfId                          string        `json:"nfId"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	Pei                           string        `json:"pei,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	Supi                          string        `json:"supi,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
	GroupId                       string        `json:"groupId,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
}
