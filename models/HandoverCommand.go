/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:16 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverCommand struct {
	Sessions              []N2SmInfoDownlinkContent `json:"sessions,omitempty"`
	NasSecFromNGRan       []byte                    `json:"nasSecFromNGRan,omitempty"`
	TargetToSourceContent []byte                    `json:"targetToSourceContent"`
	HandoverType          int16                     `json:"handoverType"`
}
