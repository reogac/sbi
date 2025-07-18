/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:38 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UserLocation struct {
	UtraLocation  *UtraLocation  `json:"utraLocation,omitempty"`
	GeraLocation  *GeraLocation  `json:"geraLocation,omitempty"`
	EutraLocation *EutraLocation `json:"eutraLocation,omitempty"`
	NrLocation    *NrLocation    `json:"nrLocation,omitempty"`
	N3gaLocation  *N3gaLocation  `json:"n3gaLocation,omitempty"`
}
