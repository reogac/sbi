/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type WlanPerformanceReq struct {
	WlanOrderCriter WlanOrderingCriterion `json:"wlanOrderCriter,omitempty"`
	Order           MatchingDirection     `json:"order,omitempty"`
	SsIds           []string              `json:"ssIds,omitempty"`
	BssIds          []string              `json:"bssIds,omitempty"`
}
