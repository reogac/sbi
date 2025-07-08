/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextReleaseRequest struct {
	Cause       N2Cause `json:"cause"`
	SessionList []int16 `json:"sessionList,omitempty"`
}
