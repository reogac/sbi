/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Mon Apr 21 15:01:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AllowedSnssai struct {
	AllowedSnssai      Snssai           `json:"allowedSnssai"`
	NsiInformationList []NsiInformation `json:"nsiInformationList,omitempty"`
	MappedHomeSnssai   *Snssai          `json:"mappedHomeSnssai,omitempty"`
}
