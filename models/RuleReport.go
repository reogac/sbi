/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RuleReport struct {
	RuleStatus      RuleStatus       `json:"ruleStatus"`
	ContVers        []int            `json:"contVers,omitempty"`
	FailureCode     FailureCode      `json:"failureCode,omitempty"`
	FinUnitAct      FinalUnitAction  `json:"finUnitAct,omitempty"`
	RanNasRelCauses []RanNasRelCause `json:"ranNasRelCauses,omitempty"`
	AltQosParamId   string           `json:"altQosParamId,omitempty"`
	PccRuleIds      []string         `json:"pccRuleIds"`
}
