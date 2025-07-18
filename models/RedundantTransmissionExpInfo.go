/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RedundantTransmissionExpInfo struct {
	SpatialValidCon *NetworkAreaInfo                `json:"spatialValidCon,omitempty"`
	Dnn             string                          `json:"dnn,omitempty"`
	RedTransExps    []RedundantTransmissionExpPerTS `json:"redTransExps"`
}
