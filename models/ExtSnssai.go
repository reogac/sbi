/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtSnssai struct {
	SdRanges   []SdRange   `json:"sdRanges,omitempty"`
	WildcardSd *WildcardSd `json:"wildcardSd,omitempty"`
	Sst        int         `json:"sst"`
	Sd         string      `json:"sd,omitempty"`
}
