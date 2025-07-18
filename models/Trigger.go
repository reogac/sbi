/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Trigger struct {
	VolumeLimit      *int            `json:"volumeLimit,omitempty"`
	VolumeLimit64    *int            `json:"volumeLimit64,omitempty"`
	EventLimit       *int            `json:"eventLimit,omitempty"`
	MaxNumberOfccc   *int            `json:"maxNumberOfccc,omitempty"`
	TariffTimeChange string          `json:"tariffTimeChange,omitempty"`
	TriggerType      TriggerType     `json:"triggerType,omitempty"`
	TriggerCategory  TriggerCategory `json:"triggerCategory"`
	TimeLimit        *int            `json:"timeLimit,omitempty"`
}
