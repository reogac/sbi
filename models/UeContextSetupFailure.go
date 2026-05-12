/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:33 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextSetupFailure struct {
	Cause      N2Cause        `json:"cause"`
	FailedList []int16        `json:"failedList,omitempty"`
	Error      ProblemDetails `json:"error"`
}
