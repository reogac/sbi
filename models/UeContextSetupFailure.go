/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:18 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextSetupFailure struct {
	FailedList []int16        `json:"failedList,omitempty"`
	Error      ProblemDetails `json:"error"`
	Cause      N2Cause        `json:"cause"`
}
