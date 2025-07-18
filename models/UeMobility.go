/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeMobility struct {
	Duration         *int                        `json:"duration,omitempty"`
	DurationVariance *float64                    `json:"durationVariance,omitempty"`
	LocInfos         []LocationInfo              `json:"locInfos,omitempty"`
	Ts               string                      `json:"ts,omitempty"`
	RecurringTime    *ScheduledCommunicationTime `json:"recurringTime,omitempty"`
}
