/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeDifferentiationInfo struct {
	BatteryInd       *BatteryIndication             `json:"batteryInd,omitempty"`
	ValidityTime     string                         `json:"validityTime,omitempty"`
	PeriodicComInd   PeriodicCommunicationIndicator `json:"periodicComInd,omitempty"`
	PeriodicTime     *int                           `json:"periodicTime,omitempty"`
	ScheduledComTime *ScheduledCommunicationTime    `json:"scheduledComTime,omitempty"`
	StationaryInd    StationaryIndication           `json:"stationaryInd,omitempty"`
	TrafficProfile   TrafficProfile                 `json:"trafficProfile,omitempty"`
}
