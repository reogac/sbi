/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NsiLoadLevelInfo struct {
	NsiId                    string           `json:"nsiId,omitempty"`
	NumOfExceedLoadLevelThr  *int             `json:"numOfExceedLoadLevelThr,omitempty"`
	ExceedLoadLevelThrInd    *bool            `json:"exceedLoadLevelThrInd,omitempty"`
	TimePeriod               *TimeWindow      `json:"timePeriod,omitempty"`
	ResUsgThrCrossTimePeriod []TimeWindow     `json:"resUsgThrCrossTimePeriod,omitempty"`
	NumOfUes                 *NumberAverage   `json:"numOfUes,omitempty"`
	NumOfPduSess             *NumberAverage   `json:"numOfPduSess,omitempty"`
	Snssai                   Snssai           `json:"snssai"`
	Confidence               *int             `json:"confidence,omitempty"`
	ResUsage                 *ResourceUsage   `json:"resUsage,omitempty"`
	NetworkArea              *NetworkAreaInfo `json:"networkArea,omitempty"`
	LoadLevelInformation     int              `json:"loadLevelInformation"`
}
