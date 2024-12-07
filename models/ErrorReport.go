/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:30 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ErrorReport struct {
	Error                *ProblemDetails     `json:"error,omitempty"`
	RuleReports          []RuleReport        `json:"ruleReports,omitempty"`
	SessRuleReports      []SessionRuleReport `json:"sessRuleReports,omitempty"`
	PolDecFailureReports []string            `json:"polDecFailureReports,omitempty"`
	InvalidPolicyDecs    []InvalidParam      `json:"invalidPolicyDecs,omitempty"`
}
