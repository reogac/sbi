/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:47 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type BdtData struct {
	Dnn         string           `json:"dnn,omitempty"`
	TrafficDes  string           `json:"trafficDes,omitempty"`
	SuppFeat    string           `json:"suppFeat,omitempty"`
	ResetIds    []string         `json:"resetIds,omitempty"`
	AspId       string           `json:"aspId"`
	VolPerUe    *UsageThreshold  `json:"volPerUe,omitempty"`
	NwAreaInfo  *NetworkAreaInfo `json:"nwAreaInfo,omitempty"`
	NumOfUes    *int             `json:"numOfUes,omitempty"`
	Snssai      *Snssai          `json:"snssai,omitempty"`
	BdtpStatus  BdtPolicyStatus  `json:"bdtpStatus,omitempty"`
	TransPolicy TransferPolicy   `json:"transPolicy"`
	BdtRefId    string           `json:"bdtRefId,omitempty"`
}
