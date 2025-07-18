/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:43 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DataRestorationNotification struct {
	RecoveryTime        string          `json:"recoveryTime,omitempty"`
	PlmnId              *PlmnId         `json:"plmnId,omitempty"`
	SupiRanges          []SupiRange     `json:"supiRanges,omitempty"`
	SNssaiList          []Snssai        `json:"sNssaiList,omitempty"`
	DnnList             []string        `json:"dnnList,omitempty"`
	UdmGroupId          string          `json:"udmGroupId,omitempty"`
	LastReplicationTime string          `json:"lastReplicationTime,omitempty"`
	GpsiRanges          []IdentityRange `json:"gpsiRanges,omitempty"`
	ResetIds            []string        `json:"resetIds,omitempty"`
}
