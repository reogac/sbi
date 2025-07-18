/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SessionRule struct {
	AuthSessAmbr *Ambr                 `json:"authSessAmbr,omitempty"`
	AuthDefQos   *AuthorizedDefaultQos `json:"authDefQos,omitempty"`
	SessRuleId   string                `json:"sessRuleId"`
	RefUmData    string                `json:"refUmData,omitempty"`
	RefUmN3gData string                `json:"refUmN3gData,omitempty"`
	RefCondData  string                `json:"refCondData,omitempty"`
}
