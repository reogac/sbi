/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextCreateError struct {
	Error                     ProblemDetails `json:"error"`
	NgapCause                 *NgApCause     `json:"ngapCause,omitempty"`
	TargetToSourceFailureData *N2InfoContent `json:"targetToSourceFailureData,omitempty"`
}
