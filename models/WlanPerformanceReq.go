/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type WlanPerformanceReq struct {
	SsIds           []string              `json:"ssIds,omitempty"`
	BssIds          []string              `json:"bssIds,omitempty"`
	WlanOrderCriter WlanOrderingCriterion `json:"wlanOrderCriter,omitempty"`
	Order           MatchingDirection     `json:"order,omitempty"`
}
