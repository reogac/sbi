/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:50 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type BdtData struct {
	BdtRefId    string           `json:"bdtRefId,omitempty"`
	VolPerUe    *UsageThreshold  `json:"volPerUe,omitempty"`
	Dnn         string           `json:"dnn,omitempty"`
	Snssai      *Snssai          `json:"snssai,omitempty"`
	BdtpStatus  BdtPolicyStatus  `json:"bdtpStatus,omitempty"`
	ResetIds    []string         `json:"resetIds,omitempty"`
	AspId       string           `json:"aspId"`
	NwAreaInfo  *NetworkAreaInfo `json:"nwAreaInfo,omitempty"`
	NumOfUes    *int             `json:"numOfUes,omitempty"`
	TrafficDes  string           `json:"trafficDes,omitempty"`
	SuppFeat    string           `json:"suppFeat,omitempty"`
	TransPolicy TransferPolicy   `json:"transPolicy"`
}
