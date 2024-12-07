/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Trigger struct {
	MaxNumberOfccc   *int            `json:"maxNumberOfccc,omitempty"`
	TariffTimeChange string          `json:"tariffTimeChange,omitempty"`
	TriggerType      TriggerType     `json:"triggerType,omitempty"`
	TriggerCategory  TriggerCategory `json:"triggerCategory"`
	TimeLimit        *int            `json:"timeLimit,omitempty"`
	VolumeLimit      *int            `json:"volumeLimit,omitempty"`
	VolumeLimit64    *int            `json:"volumeLimit64,omitempty"`
	EventLimit       *int            `json:"eventLimit,omitempty"`
}
