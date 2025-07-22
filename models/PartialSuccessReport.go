/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:34 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PartialSuccessReport struct {
	SessRuleReports         []SessionRuleReport `json:"sessRuleReports,omitempty"`
	UeCampingRep            *UeCampingRep       `json:"ueCampingRep,omitempty"`
	PolicyDecFailureReports []string            `json:"policyDecFailureReports,omitempty"`
	InvalidPolicyDecs       []InvalidParam      `json:"invalidPolicyDecs,omitempty"`
	FailureCause            FailureCause        `json:"failureCause"`
	RuleReports             []RuleReport        `json:"ruleReports,omitempty"`
}
