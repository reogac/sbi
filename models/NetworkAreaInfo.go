/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:47 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NetworkAreaInfo struct {
	Ecgis       []Ecgi            `json:"ecgis,omitempty"`
	Ncgis       []Ncgi            `json:"ncgis,omitempty"`
	GRanNodeIds []GlobalRanNodeId `json:"gRanNodeIds,omitempty"`
	Tais        []Tai             `json:"tais,omitempty"`
}
