/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NetworkPerfRequirement struct {
	RelativeRatio *int            `json:"relativeRatio,omitempty"`
	AbsoluteNum   *int            `json:"absoluteNum,omitempty"`
	NwPerfType    NetworkPerfType `json:"nwPerfType"`
}
