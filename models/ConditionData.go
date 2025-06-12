/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ConditionData struct {
	ActivationTime   string     `json:"activationTime,omitempty"`
	DeactivationTime string     `json:"deactivationTime,omitempty"`
	AccessType       AccessType `json:"accessType,omitempty"`
	RatType          RatType    `json:"ratType,omitempty"`
	CondId           string     `json:"condId"`
}
