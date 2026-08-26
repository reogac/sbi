/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 10:02:44 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type MappingOfSnssai struct {
	ServingSnssai Snssai `json:"servingSnssai"`
	HomeSnssai    Snssai `json:"homeSnssai"`
	IsLbo         bool   `json:"isLbo"`
}
