/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:15 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverCommand struct {
	Sessions              []N2SmInfoDownlinkContent `json:"sessions,omitempty"`
	TargetToSourceContent []byte                    `json:"targetToSourceContent,omitempty"`
	HandoverType          int16                     `json:"handoverType"`
}
