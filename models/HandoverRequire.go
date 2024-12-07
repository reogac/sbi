/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverRequire struct {
	HandoverType          int16                   `json:"handoverType"`
	RanInfo               EndpointInfo            `json:"ranInfo"`
	TargetId              GlobalRanNodeId         `json:"targetId"`
	Sessions              []N2SmInfoUplinkContent `json:"sessions,omitempty"`
	Cause                 N2Cause                 `json:"cause"`
	SourceToTargetContent []byte                  `json:"sourceToTargetContent,omitempty"`
}
