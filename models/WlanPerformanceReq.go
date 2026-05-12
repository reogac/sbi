/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type WlanPerformanceReq struct {
	WlanOrderCriter WlanOrderingCriterion `json:"wlanOrderCriter,omitempty"`
	Order           MatchingDirection     `json:"order,omitempty"`
	SsIds           []string              `json:"ssIds,omitempty"`
	BssIds          []string              `json:"bssIds,omitempty"`
}
