/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtSnssai struct {
	WildcardSd *WildcardSd `json:"wildcardSd,omitempty"`
	Sst        int         `json:"sst"`
	Sd         string      `json:"sd,omitempty"`
	SdRanges   []SdRange   `json:"sdRanges,omitempty"`
}
