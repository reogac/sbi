/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:34 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UePolicySetPatch struct {
	Pei              string                     `json:"pei,omitempty"`
	OsIds            []string                   `json:"osIds,omitempty"`
	UePolicySections map[string]UePolicySection `json:"uePolicySections,omitempty"`
	Upsis            []string                   `json:"upsis,omitempty"`
	AndspInd         *bool                      `json:"andspInd,omitempty"`
}
