/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:50 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDataPatch struct {
	SmPolicySnssaiData map[string]SmPolicySnssaiDataPatch `json:"smPolicySnssaiData,omitempty"`
	UmData             map[string]UsageMonData            `json:"umData,omitempty"`
}
