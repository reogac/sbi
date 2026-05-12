/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:42 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionRuleReport struct {
	RuleIds                 []string               `json:"ruleIds"`
	RuleStatus              RuleStatus             `json:"ruleStatus"`
	SessRuleFailureCode     SessionRuleFailureCode `json:"sessRuleFailureCode,omitempty"`
	PolicyDecFailureReports []string               `json:"policyDecFailureReports,omitempty"`
}
