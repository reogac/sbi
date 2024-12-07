/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:34 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyDataForIndividualUe struct {
	UePolicyDataSet         *UePolicySet                             `json:"uePolicyDataSet,omitempty"`
	SmPolicyDataSet         *SmPolicyData                            `json:"smPolicyDataSet,omitempty"`
	AmPolicyDataSet         *AmPolicyData                            `json:"amPolicyDataSet,omitempty"`
	UmData                  map[string]UsageMonData                  `json:"umData,omitempty"`
	OperatorSpecificDataSet map[string]OperatorSpecificDataContainer `json:"operatorSpecificDataSet,omitempty"`
}
