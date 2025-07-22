/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:17 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeDifferentiationInfo struct {
	PeriodicComInd   PeriodicCommunicationIndicator `json:"periodicComInd,omitempty"`
	PeriodicTime     *int                           `json:"periodicTime,omitempty"`
	ScheduledComTime *ScheduledCommunicationTime    `json:"scheduledComTime,omitempty"`
	StationaryInd    StationaryIndication           `json:"stationaryInd,omitempty"`
	TrafficProfile   TrafficProfile                 `json:"trafficProfile,omitempty"`
	BatteryInd       *BatteryIndication             `json:"batteryInd,omitempty"`
	ValidityTime     string                         `json:"validityTime,omitempty"`
}
