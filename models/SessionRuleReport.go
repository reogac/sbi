/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionRuleReport struct {
	RuleIds                 []string               `json:"ruleIds"`
	RuleStatus              RuleStatus             `json:"ruleStatus"`
	SessRuleFailureCode     SessionRuleFailureCode `json:"sessRuleFailureCode,omitempty"`
	PolicyDecFailureReports []string               `json:"policyDecFailureReports,omitempty"`
}
