/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:22 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextReleaseRequest struct {
	SessionList []int16 `json:"sessionList,omitempty"`
	Cause       N2Cause `json:"cause"`
}
