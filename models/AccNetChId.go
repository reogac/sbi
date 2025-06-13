/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AccNetChId struct {
	AccNetChaIdValue *int     `json:"accNetChaIdValue,omitempty"`
	AccNetChargId    string   `json:"accNetChargId,omitempty"`
	RefPccRuleIds    []string `json:"refPccRuleIds,omitempty"`
	SessionChScope   *bool    `json:"sessionChScope,omitempty"`
}
