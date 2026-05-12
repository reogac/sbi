/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:35 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AllowedSnssai struct {
	MappedHomeSnssai   *Snssai          `json:"mappedHomeSnssai,omitempty"`
	AllowedSnssai      Snssai           `json:"allowedSnssai"`
	NsiInformationList []NsiInformation `json:"nsiInformationList,omitempty"`
}
