/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:30 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccNetChId struct {
	RefPccRuleIds    []string `json:"refPccRuleIds,omitempty"`
	SessionChScope   *bool    `json:"sessionChScope,omitempty"`
	AccNetChaIdValue *int     `json:"accNetChaIdValue,omitempty"`
	AccNetChargId    string   `json:"accNetChargId,omitempty"`
}
