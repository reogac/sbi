/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PartialSuccessReport struct {
	FailureCause            FailureCause        `json:"failureCause"`
	RuleReports             []RuleReport        `json:"ruleReports,omitempty"`
	SessRuleReports         []SessionRuleReport `json:"sessRuleReports,omitempty"`
	UeCampingRep            *UeCampingRep       `json:"ueCampingRep,omitempty"`
	PolicyDecFailureReports []string            `json:"policyDecFailureReports,omitempty"`
	InvalidPolicyDecs       []InvalidParam      `json:"invalidPolicyDecs,omitempty"`
}
