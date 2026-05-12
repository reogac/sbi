/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:42 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RuleReport struct {
	PccRuleIds      []string         `json:"pccRuleIds"`
	RuleStatus      RuleStatus       `json:"ruleStatus"`
	ContVers        []int            `json:"contVers,omitempty"`
	FailureCode     FailureCode      `json:"failureCode,omitempty"`
	FinUnitAct      FinalUnitAction  `json:"finUnitAct,omitempty"`
	RanNasRelCauses []RanNasRelCause `json:"ranNasRelCauses,omitempty"`
	AltQosParamId   string           `json:"altQosParamId,omitempty"`
}
