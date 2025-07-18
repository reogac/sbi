/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PduSessionCreateError struct {
	RecoveryTime string           `json:"recoveryTime,omitempty"`
	Error        ProblemDetails   `json:"error"`
	N1smCause    string           `json:"n1smCause,omitempty"`
	N1SmInfoToUe *RefToBinaryData `json:"n1SmInfoToUe,omitempty"`
	BackOffTimer *int             `json:"backOffTimer,omitempty"`
}
