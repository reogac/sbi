/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:33 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UePolicySetPatch struct {
	OsIds            []string                   `json:"osIds,omitempty"`
	UePolicySections map[string]UePolicySection `json:"uePolicySections,omitempty"`
	Upsis            []string                   `json:"upsis,omitempty"`
	AndspInd         *bool                      `json:"andspInd,omitempty"`
	Pei              string                     `json:"pei,omitempty"`
}
