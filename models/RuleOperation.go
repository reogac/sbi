/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RuleOperation string

// Define constant values for RuleOperation
const (
	RULEOPERATION_CREATE_PCC_RULE                               RuleOperation = "CREATE_PCC_RULE"
	RULEOPERATION_DELETE_PCC_RULE                               RuleOperation = "DELETE_PCC_RULE"
	RULEOPERATION_MODIFY_PCC_RULE_AND_ADD_PACKET_FILTERS        RuleOperation = "MODIFY_PCC_RULE_AND_ADD_PACKET_FILTERS"
	RULEOPERATION_MODIFY_PCC_RULE_AND_REPLACE_PACKET_FILTERS    RuleOperation = "MODIFY_PCC_RULE_AND_REPLACE_PACKET_FILTERS"
	RULEOPERATION_MODIFY_PCC_RULE_AND_DELETE_PACKET_FILTERS     RuleOperation = "MODIFY_PCC_RULE_AND_DELETE_PACKET_FILTERS"
	RULEOPERATION_MODIFY_PCC_RULE_WITHOUT_MODIFY_PACKET_FILTERS RuleOperation = "MODIFY_PCC_RULE_WITHOUT_MODIFY_PACKET_FILTERS"
)
