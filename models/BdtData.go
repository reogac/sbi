/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type BdtData struct {
	NwAreaInfo  *NetworkAreaInfo `json:"nwAreaInfo,omitempty"`
	Dnn         string           `json:"dnn,omitempty"`
	Snssai      *Snssai          `json:"snssai,omitempty"`
	ResetIds    []string         `json:"resetIds,omitempty"`
	AspId       string           `json:"aspId"`
	TransPolicy TransferPolicy   `json:"transPolicy"`
	BdtRefId    string           `json:"bdtRefId,omitempty"`
	BdtpStatus  BdtPolicyStatus  `json:"bdtpStatus,omitempty"`
	SuppFeat    string           `json:"suppFeat,omitempty"`
	NumOfUes    *int             `json:"numOfUes,omitempty"`
	VolPerUe    *UsageThreshold  `json:"volPerUe,omitempty"`
	TrafficDes  string           `json:"trafficDes,omitempty"`
}
