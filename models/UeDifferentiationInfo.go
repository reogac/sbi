/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeDifferentiationInfo struct {
	ScheduledComTime *ScheduledCommunicationTime    `json:"scheduledComTime,omitempty"`
	StationaryInd    StationaryIndication           `json:"stationaryInd,omitempty"`
	TrafficProfile   TrafficProfile                 `json:"trafficProfile,omitempty"`
	BatteryInd       *BatteryIndication             `json:"batteryInd,omitempty"`
	ValidityTime     string                         `json:"validityTime,omitempty"`
	PeriodicComInd   PeriodicCommunicationIndicator `json:"periodicComInd,omitempty"`
	PeriodicTime     *int                           `json:"periodicTime,omitempty"`
}
