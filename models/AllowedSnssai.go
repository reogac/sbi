/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:24 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AllowedSnssai struct {
	MappedHomeSnssai   *Snssai          `json:"mappedHomeSnssai,omitempty"`
	AllowedSnssai      Snssai           `json:"allowedSnssai"`
	NsiInformationList []NsiInformation `json:"nsiInformationList,omitempty"`
}
