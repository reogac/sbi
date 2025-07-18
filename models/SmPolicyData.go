/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyData struct {
	UmDataLimits       map[string]UsageMonDataLimit  `json:"umDataLimits,omitempty"`
	UmData             map[string]UsageMonData       `json:"umData,omitempty"`
	SuppFeat           string                        `json:"suppFeat,omitempty"`
	ResetIds           []string                      `json:"resetIds,omitempty"`
	SmPolicySnssaiData map[string]SmPolicySnssaiData `json:"smPolicySnssaiData"`
}
