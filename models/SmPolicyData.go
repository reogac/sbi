/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:46 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyData struct {
	UmData             map[string]UsageMonData       `json:"umData,omitempty"`
	SuppFeat           string                        `json:"suppFeat,omitempty"`
	ResetIds           []string                      `json:"resetIds,omitempty"`
	SmPolicySnssaiData map[string]SmPolicySnssaiData `json:"smPolicySnssaiData"`
	UmDataLimits       map[string]UsageMonDataLimit  `json:"umDataLimits,omitempty"`
}
