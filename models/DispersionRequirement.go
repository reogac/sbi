/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DispersionRequirement struct {
	DisperType      DispersionType              `json:"disperType"`
	ClassCriters    []ClassCriterion            `json:"classCriters,omitempty"`
	RankCriters     []RankingCriterion          `json:"rankCriters,omitempty"`
	DispOrderCriter DispersionOrderingCriterion `json:"dispOrderCriter,omitempty"`
	Order           MatchingDirection           `json:"order,omitempty"`
}
