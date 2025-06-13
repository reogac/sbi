/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:28 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UserLocation struct {
	EutraLocation *EutraLocation `json:"eutraLocation,omitempty"`
	NrLocation    *NrLocation    `json:"nrLocation,omitempty"`
	N3gaLocation  *N3gaLocation  `json:"n3gaLocation,omitempty"`
	UtraLocation  *UtraLocation  `json:"utraLocation,omitempty"`
	GeraLocation  *GeraLocation  `json:"geraLocation,omitempty"`
}
