/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:52 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDataPatch struct {
	SmPolicySnssaiData map[string]SmPolicySnssaiDataPatch `json:"smPolicySnssaiData,omitempty"`
	UmData             map[string]UsageMonData            `json:"umData,omitempty"`
}
