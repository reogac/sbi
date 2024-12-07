/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:17 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UplinkNasError struct {
	CauseValue   int16  `json:"causeValue"`
	NasPdu       []byte `json:"nasPdu"`
	CausePresent int16  `json:"causePresent"`
}
