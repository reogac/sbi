/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:22 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextSetupFailure struct {
	Error        ProblemDetails `json:"error"`
	CausePresent int16          `json:"causePresent"`
	CauseValue   int16          `json:"causeValue"`
	FailedList   []int16        `json:"failedList,omitempty"`
}
