/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:47:48 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextReleaseRequest struct {
	Cause       N2Cause `json:"cause"`
	SessionList []int16 `json:"sessionList,omitempty"`
}
