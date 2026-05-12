/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TargetUeInformation struct {
	AnyUe       *bool    `json:"anyUe,omitempty"`
	Supis       []string `json:"supis,omitempty"`
	Gpsis       []string `json:"gpsis,omitempty"`
	IntGroupIds []string `json:"intGroupIds,omitempty"`
}
