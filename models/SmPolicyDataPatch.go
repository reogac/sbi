/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:46 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDataPatch struct {
	UmData             map[string]UsageMonData            `json:"umData,omitempty"`
	SmPolicySnssaiData map[string]SmPolicySnssaiDataPatch `json:"smPolicySnssaiData,omitempty"`
}
