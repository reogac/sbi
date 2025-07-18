/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NetworkPerfInfo struct {
	Confidence    *int             `json:"confidence,omitempty"`
	NetworkArea   *NetworkAreaInfo `json:"networkArea,omitempty"`
	NwPerfType    NetworkPerfType  `json:"nwPerfType,omitempty"`
	RelativeRatio *int             `json:"relativeRatio,omitempty"`
	AbsoluteNum   *int             `json:"absoluteNum,omitempty"`
}
