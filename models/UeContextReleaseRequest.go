/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:43 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextReleaseRequest struct {
	SessionList []int16 `json:"sessionList,omitempty"`
	Cause       N2Cause `json:"cause"`
}
