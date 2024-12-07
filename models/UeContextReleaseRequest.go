/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:16 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextReleaseRequest struct {
	SessionList  []int16 `json:"sessionList,omitempty"`
	CausePresent int16   `json:"causePresent"`
	CauseValue   int16   `json:"causeValue"`
}
