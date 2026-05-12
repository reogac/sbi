/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:42 KST 2026 by TungTQ<tqtung@etri.re.kr>
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
