/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDataPatch struct {
	UmData             map[string]UsageMonData            `json:"umData,omitempty"`
	SmPolicySnssaiData map[string]SmPolicySnssaiDataPatch `json:"smPolicySnssaiData,omitempty"`
}
