/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:30 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TargetUeInformation struct {
	Gpsis       []string `json:"gpsis,omitempty"`
	IntGroupIds []string `json:"intGroupIds,omitempty"`
	AnyUe       *bool    `json:"anyUe,omitempty"`
	Supis       []string `json:"supis,omitempty"`
}
