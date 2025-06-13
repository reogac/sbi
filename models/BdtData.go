/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type BdtData struct {
	Dnn         string           `json:"dnn,omitempty"`
	TrafficDes  string           `json:"trafficDes,omitempty"`
	BdtpStatus  BdtPolicyStatus  `json:"bdtpStatus,omitempty"`
	SuppFeat    string           `json:"suppFeat,omitempty"`
	AspId       string           `json:"aspId"`
	BdtRefId    string           `json:"bdtRefId,omitempty"`
	NwAreaInfo  *NetworkAreaInfo `json:"nwAreaInfo,omitempty"`
	Snssai      *Snssai          `json:"snssai,omitempty"`
	ResetIds    []string         `json:"resetIds,omitempty"`
	TransPolicy TransferPolicy   `json:"transPolicy"`
	NumOfUes    *int             `json:"numOfUes,omitempty"`
	VolPerUe    *UsageThreshold  `json:"volPerUe,omitempty"`
}
