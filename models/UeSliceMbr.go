/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeSliceMbr struct {
	SliceMbr         map[string]SliceMbr `json:"sliceMbr"`
	ServingSnssai    Snssai              `json:"servingSnssai"`
	MappedHomeSnssai *Snssai             `json:"mappedHomeSnssai,omitempty"`
}
