/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type LocationArea struct {
	GeographicAreas []GeographicArea `json:"geographicAreas,omitempty"`
	CivicAddresses  []CivicAddress   `json:"civicAddresses,omitempty"`
	NwAreaInfo      *NetworkAreaInfo `json:"nwAreaInfo,omitempty"`
	UmtTime         *UmtTime         `json:"umtTime,omitempty"`
}
