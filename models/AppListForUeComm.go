/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:27 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AppListForUeComm struct {
	SpatialValidity *NetworkAreaInfo `json:"spatialValidity,omitempty"`
	AppId           string           `json:"appId"`
	StartTime       string           `json:"startTime,omitempty"`
	AppDur          *int             `json:"appDur,omitempty"`
	OccurRatio      *int             `json:"occurRatio,omitempty"`
}
