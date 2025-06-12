/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NetworkAreaInfo struct {
	Ncgis       []Ncgi            `json:"ncgis,omitempty"`
	GRanNodeIds []GlobalRanNodeId `json:"gRanNodeIds,omitempty"`
	Tais        []Tai             `json:"tais,omitempty"`
	Ecgis       []Ecgi            `json:"ecgis,omitempty"`
}
