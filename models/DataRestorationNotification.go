/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:23 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DataRestorationNotification struct {
	LastReplicationTime string          `json:"lastReplicationTime,omitempty"`
	PlmnId              *PlmnId         `json:"plmnId,omitempty"`
	UdmGroupId          string          `json:"udmGroupId,omitempty"`
	RecoveryTime        string          `json:"recoveryTime,omitempty"`
	SupiRanges          []SupiRange     `json:"supiRanges,omitempty"`
	GpsiRanges          []IdentityRange `json:"gpsiRanges,omitempty"`
	ResetIds            []string        `json:"resetIds,omitempty"`
	SNssaiList          []Snssai        `json:"sNssaiList,omitempty"`
	DnnList             []string        `json:"dnnList,omitempty"`
}
