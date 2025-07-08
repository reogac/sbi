/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HsmfUpdateError struct {
	RecoveryTime string           `json:"recoveryTime,omitempty"`
	Error        ProblemDetails   `json:"error"`
	Pti          *int             `json:"pti,omitempty"`
	N1smCause    string           `json:"n1smCause,omitempty"`
	N1SmInfoToUe *RefToBinaryData `json:"n1SmInfoToUe,omitempty"`
	BackOffTimer *int             `json:"backOffTimer,omitempty"`
}
