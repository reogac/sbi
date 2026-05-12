/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NsiLoadLevelInfo struct {
	TimePeriod               *TimeWindow      `json:"timePeriod,omitempty"`
	ResUsgThrCrossTimePeriod []TimeWindow     `json:"resUsgThrCrossTimePeriod,omitempty"`
	NumOfPduSess             *NumberAverage   `json:"numOfPduSess,omitempty"`
	Snssai                   Snssai           `json:"snssai"`
	ResUsage                 *ResourceUsage   `json:"resUsage,omitempty"`
	NumOfUes                 *NumberAverage   `json:"numOfUes,omitempty"`
	Confidence               *int             `json:"confidence,omitempty"`
	LoadLevelInformation     int              `json:"loadLevelInformation"`
	NsiId                    string           `json:"nsiId,omitempty"`
	NumOfExceedLoadLevelThr  *int             `json:"numOfExceedLoadLevelThr,omitempty"`
	ExceedLoadLevelThrInd    *bool            `json:"exceedLoadLevelThrInd,omitempty"`
	NetworkArea              *NetworkAreaInfo `json:"networkArea,omitempty"`
}
