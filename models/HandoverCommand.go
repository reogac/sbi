/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:21 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverCommand struct {
	TargetToSourceContent []byte                    `json:"targetToSourceContent"`
	HandoverType          int16                     `json:"handoverType"`
	Sessions              []N2SmInfoDownlinkContent `json:"sessions,omitempty"`
	NasSecFromNGRan       []byte                    `json:"nasSecFromNGRan,omitempty"`
}
