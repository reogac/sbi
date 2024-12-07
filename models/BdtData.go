/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:34 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type BdtData struct {
	Dnn         string           `json:"dnn,omitempty"`
	TransPolicy TransferPolicy   `json:"transPolicy"`
	BdtRefId    string           `json:"bdtRefId,omitempty"`
	NumOfUes    *int             `json:"numOfUes,omitempty"`
	Snssai      *Snssai          `json:"snssai,omitempty"`
	TrafficDes  string           `json:"trafficDes,omitempty"`
	BdtpStatus  BdtPolicyStatus  `json:"bdtpStatus,omitempty"`
	SuppFeat    string           `json:"suppFeat,omitempty"`
	ResetIds    []string         `json:"resetIds,omitempty"`
	AspId       string           `json:"aspId"`
	NwAreaInfo  *NetworkAreaInfo `json:"nwAreaInfo,omitempty"`
	VolPerUe    *UsageThreshold  `json:"volPerUe,omitempty"`
}
