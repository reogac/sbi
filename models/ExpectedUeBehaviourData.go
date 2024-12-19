/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:49:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExpectedUeBehaviourData struct {
	ValidityTime               string                      `json:"validityTime,omitempty"`
	StationaryIndication       StationaryIndication        `json:"stationaryIndication,omitempty"`
	PeriodicTime               *int                        `json:"periodicTime,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`
	ScheduledCommunicationType ScheduledCommunicationType  `json:"scheduledCommunicationType,omitempty"`
	BatteryIndication          *BatteryIndication          `json:"batteryIndication,omitempty"`
	CommunicationDurationTime  *int                        `json:"communicationDurationTime,omitempty"`
	ExpectedUmts               []LocationArea              `json:"expectedUmts,omitempty"`
	TrafficProfile             TrafficProfile              `json:"trafficProfile,omitempty"`
}
