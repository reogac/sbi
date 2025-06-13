/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type BdtData struct {
	TrafficDes  string           `json:"trafficDes,omitempty"`
	BdtpStatus  BdtPolicyStatus  `json:"bdtpStatus,omitempty"`
	SuppFeat    string           `json:"suppFeat,omitempty"`
	NwAreaInfo  *NetworkAreaInfo `json:"nwAreaInfo,omitempty"`
	VolPerUe    *UsageThreshold  `json:"volPerUe,omitempty"`
	Dnn         string           `json:"dnn,omitempty"`
	NumOfUes    *int             `json:"numOfUes,omitempty"`
	Snssai      *Snssai          `json:"snssai,omitempty"`
	ResetIds    []string         `json:"resetIds,omitempty"`
	AspId       string           `json:"aspId"`
	TransPolicy TransferPolicy   `json:"transPolicy"`
	BdtRefId    string           `json:"bdtRefId,omitempty"`
}
