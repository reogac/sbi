/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExtSnssai struct {
	WildcardSd *WildcardSd `json:"wildcardSd,omitempty"`
	Sd         string      `json:"sd,omitempty"`
	Sst        int         `json:"sst"`
	SdRanges   []SdRange   `json:"sdRanges,omitempty"`
}
