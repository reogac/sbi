/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HsmfUpdateError struct {
	BackOffTimer *int             `json:"backOffTimer,omitempty"`
	RecoveryTime string           `json:"recoveryTime,omitempty"`
	Error        ProblemDetails   `json:"error"`
	Pti          *int             `json:"pti,omitempty"`
	N1smCause    string           `json:"n1smCause,omitempty"`
	N1SmInfoToUe *RefToBinaryData `json:"n1SmInfoToUe,omitempty"`
}
