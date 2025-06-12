/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccNetChId struct {
	AccNetChaIdValue *int     `json:"accNetChaIdValue,omitempty"`
	AccNetChargId    string   `json:"accNetChargId,omitempty"`
	RefPccRuleIds    []string `json:"refPccRuleIds,omitempty"`
	SessionChScope   *bool    `json:"sessionChScope,omitempty"`
}
