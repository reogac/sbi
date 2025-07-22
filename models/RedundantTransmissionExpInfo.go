/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RedundantTransmissionExpInfo struct {
	SpatialValidCon *NetworkAreaInfo                `json:"spatialValidCon,omitempty"`
	Dnn             string                          `json:"dnn,omitempty"`
	RedTransExps    []RedundantTransmissionExpPerTS `json:"redTransExps"`
}
