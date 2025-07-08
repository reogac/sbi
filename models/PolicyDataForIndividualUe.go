/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:47 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyDataForIndividualUe struct {
	SmPolicyDataSet         *SmPolicyData                            `json:"smPolicyDataSet,omitempty"`
	AmPolicyDataSet         *AmPolicyData                            `json:"amPolicyDataSet,omitempty"`
	UmData                  map[string]UsageMonData                  `json:"umData,omitempty"`
	OperatorSpecificDataSet map[string]OperatorSpecificDataContainer `json:"operatorSpecificDataSet,omitempty"`
	UePolicyDataSet         *UePolicySet                             `json:"uePolicyDataSet,omitempty"`
}
