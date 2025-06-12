/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeSliceMbr struct {
	ServingSnssai    Snssai              `json:"servingSnssai"`
	MappedHomeSnssai *Snssai             `json:"mappedHomeSnssai,omitempty"`
	SliceMbr         map[string]SliceMbr `json:"sliceMbr"`
}
