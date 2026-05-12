/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:39 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DataRestorationNotification struct {
	SupiRanges          []SupiRange     `json:"supiRanges,omitempty"`
	ResetIds            []string        `json:"resetIds,omitempty"`
	SNssaiList          []Snssai        `json:"sNssaiList,omitempty"`
	DnnList             []string        `json:"dnnList,omitempty"`
	UdmGroupId          string          `json:"udmGroupId,omitempty"`
	LastReplicationTime string          `json:"lastReplicationTime,omitempty"`
	RecoveryTime        string          `json:"recoveryTime,omitempty"`
	GpsiRanges          []IdentityRange `json:"gpsiRanges,omitempty"`
	PlmnId              *PlmnId         `json:"plmnId,omitempty"`
}
