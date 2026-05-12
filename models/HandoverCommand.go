/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:27 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type HandoverCommand struct {
	NasSecFromNGRan       []byte                    `json:"nasSecFromNGRan,omitempty"`
	TargetToSourceContent []byte                    `json:"targetToSourceContent"`
	HandoverType          int16                     `json:"handoverType"`
	Sessions              []N2SmInfoDownlinkContent `json:"sessions,omitempty"`
}
