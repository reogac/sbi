/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AssignEbiError struct {
	Error          ProblemDetails  `json:"error"`
	FailureDetails AssignEbiFailed `json:"failureDetails"`
}
