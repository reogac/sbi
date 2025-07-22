/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DataRestorationNotification struct {
	GpsiRanges          []IdentityRange `json:"gpsiRanges,omitempty"`
	DnnList             []string        `json:"dnnList,omitempty"`
	UdmGroupId          string          `json:"udmGroupId,omitempty"`
	RecoveryTime        string          `json:"recoveryTime,omitempty"`
	SupiRanges          []SupiRange     `json:"supiRanges,omitempty"`
	ResetIds            []string        `json:"resetIds,omitempty"`
	SNssaiList          []Snssai        `json:"sNssaiList,omitempty"`
	LastReplicationTime string          `json:"lastReplicationTime,omitempty"`
	PlmnId              *PlmnId         `json:"plmnId,omitempty"`
}
