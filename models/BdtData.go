/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type BdtData struct {
	AspId       string           `json:"aspId"`
	TransPolicy TransferPolicy   `json:"transPolicy"`
	BdtRefId    string           `json:"bdtRefId,omitempty"`
	NwAreaInfo  *NetworkAreaInfo `json:"nwAreaInfo,omitempty"`
	VolPerUe    *UsageThreshold  `json:"volPerUe,omitempty"`
	Snssai      *Snssai          `json:"snssai,omitempty"`
	TrafficDes  string           `json:"trafficDes,omitempty"`
	BdtpStatus  BdtPolicyStatus  `json:"bdtpStatus,omitempty"`
	ResetIds    []string         `json:"resetIds,omitempty"`
	NumOfUes    *int             `json:"numOfUes,omitempty"`
	Dnn         string           `json:"dnn,omitempty"`
	SuppFeat    string           `json:"suppFeat,omitempty"`
}
