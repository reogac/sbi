/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:38 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type BdtData struct {
	SuppFeat    string           `json:"suppFeat,omitempty"`
	ResetIds    []string         `json:"resetIds,omitempty"`
	BdtRefId    string           `json:"bdtRefId,omitempty"`
	NwAreaInfo  *NetworkAreaInfo `json:"nwAreaInfo,omitempty"`
	VolPerUe    *UsageThreshold  `json:"volPerUe,omitempty"`
	Snssai      *Snssai          `json:"snssai,omitempty"`
	TrafficDes  string           `json:"trafficDes,omitempty"`
	BdtpStatus  BdtPolicyStatus  `json:"bdtpStatus,omitempty"`
	AspId       string           `json:"aspId"`
	TransPolicy TransferPolicy   `json:"transPolicy"`
	NumOfUes    *int             `json:"numOfUes,omitempty"`
	Dnn         string           `json:"dnn,omitempty"`
}
