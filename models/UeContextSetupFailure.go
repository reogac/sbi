/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:28 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextSetupFailure struct {
	Cause      N2Cause        `json:"cause"`
	FailedList []int16        `json:"failedList,omitempty"`
	Error      ProblemDetails `json:"error"`
}
