/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AppListForUeComm struct {
	AppId           string           `json:"appId"`
	StartTime       string           `json:"startTime,omitempty"`
	AppDur          *int             `json:"appDur,omitempty"`
	OccurRatio      *int             `json:"occurRatio,omitempty"`
	SpatialValidity *NetworkAreaInfo `json:"spatialValidity,omitempty"`
}
