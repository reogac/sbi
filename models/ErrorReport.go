/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:42 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ErrorReport struct {
	RuleReports          []RuleReport        `json:"ruleReports,omitempty"`
	SessRuleReports      []SessionRuleReport `json:"sessRuleReports,omitempty"`
	PolDecFailureReports []string            `json:"polDecFailureReports,omitempty"`
	InvalidPolicyDecs    []InvalidParam      `json:"invalidPolicyDecs,omitempty"`
	Error                *ProblemDetails     `json:"error,omitempty"`
}
