/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:03 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyDataForIndividualUe struct {
	AmPolicyDataSet         *AmPolicyData                            `json:"amPolicyDataSet,omitempty"`
	UmData                  map[string]UsageMonData                  `json:"umData,omitempty"`
	OperatorSpecificDataSet map[string]OperatorSpecificDataContainer `json:"operatorSpecificDataSet,omitempty"`
	UePolicyDataSet         *UePolicySet                             `json:"uePolicyDataSet,omitempty"`
	SmPolicyDataSet         *SmPolicyData                            `json:"smPolicyDataSet,omitempty"`
}
