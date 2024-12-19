/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 11:08:48 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AllowedSnssai struct {
	AllowedSnssai      Snssai           `json:"allowedSnssai"`
	NsiInformationList []NsiInformation `json:"nsiInformationList,omitempty"`
	MappedHomeSnssai   *Snssai          `json:"mappedHomeSnssai,omitempty"`
}
