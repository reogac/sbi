/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:43 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeSliceMbr struct {
	SliceMbr         map[string]SliceMbr `json:"sliceMbr"`
	ServingSnssai    Snssai              `json:"servingSnssai"`
	MappedHomeSnssai *Snssai             `json:"mappedHomeSnssai,omitempty"`
}
