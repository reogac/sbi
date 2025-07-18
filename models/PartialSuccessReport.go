/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PartialSuccessReport struct {
	InvalidPolicyDecs       []InvalidParam      `json:"invalidPolicyDecs,omitempty"`
	FailureCause            FailureCause        `json:"failureCause"`
	RuleReports             []RuleReport        `json:"ruleReports,omitempty"`
	SessRuleReports         []SessionRuleReport `json:"sessRuleReports,omitempty"`
	UeCampingRep            *UeCampingRep       `json:"ueCampingRep,omitempty"`
	PolicyDecFailureReports []string            `json:"policyDecFailureReports,omitempty"`
}
