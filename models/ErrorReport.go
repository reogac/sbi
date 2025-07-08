/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:44 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
