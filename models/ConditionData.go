/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:34 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ConditionData struct {
	RatType          RatType    `json:"ratType,omitempty"`
	CondId           string     `json:"condId"`
	ActivationTime   string     `json:"activationTime,omitempty"`
	DeactivationTime string     `json:"deactivationTime,omitempty"`
	AccessType       AccessType `json:"accessType,omitempty"`
}
