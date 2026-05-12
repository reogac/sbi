/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:46 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type BdtData struct {
	TransPolicy TransferPolicy   `json:"transPolicy"`
	NwAreaInfo  *NetworkAreaInfo `json:"nwAreaInfo,omitempty"`
	NumOfUes    *int             `json:"numOfUes,omitempty"`
	Dnn         string           `json:"dnn,omitempty"`
	TrafficDes  string           `json:"trafficDes,omitempty"`
	BdtpStatus  BdtPolicyStatus  `json:"bdtpStatus,omitempty"`
	SuppFeat    string           `json:"suppFeat,omitempty"`
	ResetIds    []string         `json:"resetIds,omitempty"`
	AspId       string           `json:"aspId"`
	BdtRefId    string           `json:"bdtRefId,omitempty"`
	VolPerUe    *UsageThreshold  `json:"volPerUe,omitempty"`
	Snssai      *Snssai          `json:"snssai,omitempty"`
}
