/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExpectedUeBehaviourData struct {
	ExpectedUmts               []LocationArea              `json:"expectedUmts,omitempty"`
	TrafficProfile             TrafficProfile              `json:"trafficProfile,omitempty"`
	ValidityTime               string                      `json:"validityTime,omitempty"`
	PeriodicTime               *int                        `json:"periodicTime,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`
	ScheduledCommunicationType ScheduledCommunicationType  `json:"scheduledCommunicationType,omitempty"`
	StationaryIndication       StationaryIndication        `json:"stationaryIndication,omitempty"`
	CommunicationDurationTime  *int                        `json:"communicationDurationTime,omitempty"`
	BatteryIndication          *BatteryIndication          `json:"batteryIndication,omitempty"`
}
