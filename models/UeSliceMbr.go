/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeSliceMbr struct {
	SliceMbr         map[string]SliceMbr `json:"sliceMbr"`
	ServingSnssai    Snssai              `json:"servingSnssai"`
	MappedHomeSnssai *Snssai             `json:"mappedHomeSnssai,omitempty"`
}
