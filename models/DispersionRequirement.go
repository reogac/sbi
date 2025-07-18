/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type DispersionRequirement struct {
	RankCriters     []RankingCriterion          `json:"rankCriters,omitempty"`
	DispOrderCriter DispersionOrderingCriterion `json:"dispOrderCriter,omitempty"`
	Order           MatchingDirection           `json:"order,omitempty"`
	DisperType      DispersionType              `json:"disperType"`
	ClassCriters    []ClassCriterion            `json:"classCriters,omitempty"`
}
