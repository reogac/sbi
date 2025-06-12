/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EcsAddrConfigInfo struct {
	Snssai              *Snssai              `json:"snssai,omitempty"`
	EcsServerAddr       *EcsServerAddr       `json:"ecsServerAddr,omitempty"`
	SpatialValidityCond *SpatialValidityCond `json:"spatialValidityCond,omitempty"`
	Dnn                 string               `json:"dnn,omitempty"`
}
