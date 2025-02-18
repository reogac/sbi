/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Feb 18 15:01:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverRequired struct {
	Cause                 N2Cause                 `json:"cause"`
	DirectFwdPathFlag     *bool                   `json:"directFwdPathFlag,omitempty"`
	SourceToTargetContent []byte                  `json:"sourceToTargetContent,omitempty"`
	HandoverType          int16                   `json:"handoverType"`
	TargetId              GlobalRanNodeId         `json:"targetId"`
	Sessions              []N2SmInfoUplinkContent `json:"sessions"`
}
