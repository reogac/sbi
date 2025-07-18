/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:21 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverRequired struct {
	Sessions              []N2SmInfoUplinkContent `json:"sessions"`
	Cause                 N2Cause                 `json:"cause"`
	DirectFwdPathFlag     *bool                   `json:"directFwdPathFlag,omitempty"`
	SourceToTargetContent []byte                  `json:"sourceToTargetContent,omitempty"`
	HandoverType          int16                   `json:"handoverType"`
	TargetId              GlobalRanNodeId         `json:"targetId"`
}
