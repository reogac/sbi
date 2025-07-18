/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfEventSubscription struct {
	Pei                           string        `json:"pei,omitempty"`
	AnyUE                         *bool         `json:"anyUE,omitempty"`
	EventNotifyUri                string        `json:"eventNotifyUri"`
	NfId                          string        `json:"nfId"`
	SubsChangeNotifyUri           string        `json:"subsChangeNotifyUri,omitempty"`
	ExcludeGpsiList               []string      `json:"excludeGpsiList,omitempty"`
	SourceNfType                  NFType        `json:"sourceNfType,omitempty"`
	EventList                     []AmfEvent    `json:"eventList"`
	GroupId                       string        `json:"groupId,omitempty"`
	IncludeSupiList               []string      `json:"includeSupiList,omitempty"`
	IncludeGpsiList               []string      `json:"includeGpsiList,omitempty"`
	NotifyCorrelationId           string        `json:"notifyCorrelationId"`
	Supi                          string        `json:"supi,omitempty"`
	ExcludeSupiList               []string      `json:"excludeSupiList,omitempty"`
	Gpsi                          string        `json:"gpsi,omitempty"`
	Options                       *AmfEventMode `json:"options,omitempty"`
	SubsChangeNotifyCorrelationId string        `json:"subsChangeNotifyCorrelationId,omitempty"`
}
