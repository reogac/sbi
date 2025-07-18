/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HsmfUpdateError struct {
	Error        ProblemDetails   `json:"error"`
	Pti          *int             `json:"pti,omitempty"`
	N1smCause    string           `json:"n1smCause,omitempty"`
	N1SmInfoToUe *RefToBinaryData `json:"n1SmInfoToUe,omitempty"`
	BackOffTimer *int             `json:"backOffTimer,omitempty"`
	RecoveryTime string           `json:"recoveryTime,omitempty"`
}
