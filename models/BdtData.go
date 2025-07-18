/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type BdtData struct {
	NwAreaInfo  *NetworkAreaInfo `json:"nwAreaInfo,omitempty"`
	NumOfUes    *int             `json:"numOfUes,omitempty"`
	Snssai      *Snssai          `json:"snssai,omitempty"`
	TrafficDes  string           `json:"trafficDes,omitempty"`
	ResetIds    []string         `json:"resetIds,omitempty"`
	AspId       string           `json:"aspId"`
	BdtRefId    string           `json:"bdtRefId,omitempty"`
	Dnn         string           `json:"dnn,omitempty"`
	BdtpStatus  BdtPolicyStatus  `json:"bdtpStatus,omitempty"`
	SuppFeat    string           `json:"suppFeat,omitempty"`
	TransPolicy TransferPolicy   `json:"transPolicy"`
	VolPerUe    *UsageThreshold  `json:"volPerUe,omitempty"`
}
