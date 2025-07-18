/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextCreateError struct {
	Error        ExtProblemDetails `json:"error"`
	N1SmMsg      *RefToBinaryData  `json:"n1SmMsg,omitempty"`
	N2SmInfo     *RefToBinaryData  `json:"n2SmInfo,omitempty"`
	N2SmInfoType N2SmInfoType      `json:"n2SmInfoType,omitempty"`
	RecoveryTime string            `json:"recoveryTime,omitempty"`
}
