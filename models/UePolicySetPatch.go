/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:46 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UePolicySetPatch struct {
	UePolicySections map[string]UePolicySection `json:"uePolicySections,omitempty"`
	Upsis            []string                   `json:"upsis,omitempty"`
	AndspInd         *bool                      `json:"andspInd,omitempty"`
	Pei              string                     `json:"pei,omitempty"`
	OsIds            []string                   `json:"osIds,omitempty"`
}
