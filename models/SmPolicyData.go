/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyData struct {
	SmPolicySnssaiData map[string]SmPolicySnssaiData `json:"smPolicySnssaiData"`
	UmDataLimits       map[string]UsageMonDataLimit  `json:"umDataLimits,omitempty"`
	UmData             map[string]UsageMonData       `json:"umData,omitempty"`
	SuppFeat           string                        `json:"suppFeat,omitempty"`
	ResetIds           []string                      `json:"resetIds,omitempty"`
}
