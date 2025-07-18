/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NsiLoadLevelInfo struct {
	Snssai                   Snssai           `json:"snssai"`
	NsiId                    string           `json:"nsiId,omitempty"`
	NetworkArea              *NetworkAreaInfo `json:"networkArea,omitempty"`
	TimePeriod               *TimeWindow      `json:"timePeriod,omitempty"`
	Confidence               *int             `json:"confidence,omitempty"`
	LoadLevelInformation     int              `json:"loadLevelInformation"`
	NumOfExceedLoadLevelThr  *int             `json:"numOfExceedLoadLevelThr,omitempty"`
	ExceedLoadLevelThrInd    *bool            `json:"exceedLoadLevelThrInd,omitempty"`
	ResUsgThrCrossTimePeriod []TimeWindow     `json:"resUsgThrCrossTimePeriod,omitempty"`
	NumOfUes                 *NumberAverage   `json:"numOfUes,omitempty"`
	NumOfPduSess             *NumberAverage   `json:"numOfPduSess,omitempty"`
	ResUsage                 *ResourceUsage   `json:"resUsage,omitempty"`
}
