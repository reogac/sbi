/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionRuleReport struct {
	SessRuleFailureCode     SessionRuleFailureCode `json:"sessRuleFailureCode,omitempty"`
	PolicyDecFailureReports []string               `json:"policyDecFailureReports,omitempty"`
	RuleIds                 []string               `json:"ruleIds"`
	RuleStatus              RuleStatus             `json:"ruleStatus"`
}
