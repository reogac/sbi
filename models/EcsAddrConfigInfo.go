/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EcsAddrConfigInfo struct {
	Snssai              *Snssai              `json:"snssai,omitempty"`
	EcsServerAddr       *EcsServerAddr       `json:"ecsServerAddr,omitempty"`
	SpatialValidityCond *SpatialValidityCond `json:"spatialValidityCond,omitempty"`
	Dnn                 string               `json:"dnn,omitempty"`
}
