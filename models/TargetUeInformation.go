/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:14 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type TargetUeInformation struct {
	Supis       []string `json:"supis,omitempty"`
	Gpsis       []string `json:"gpsis,omitempty"`
	IntGroupIds []string `json:"intGroupIds,omitempty"`
	AnyUe       *bool    `json:"anyUe,omitempty"`
}
