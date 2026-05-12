/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:28 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UplinkNasError struct {
	NasPdu []byte  `json:"nasPdu"`
	Cause  N2Cause `json:"cause"`
}
