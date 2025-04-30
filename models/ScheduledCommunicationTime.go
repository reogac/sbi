/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Apr 30 14:54:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ScheduledCommunicationTime struct {
	TimeOfDayEnd   string `json:"timeOfDayEnd,omitempty"`
	DaysOfWeek     []int  `json:"daysOfWeek,omitempty"`
	TimeOfDayStart string `json:"timeOfDayStart,omitempty"`
}
