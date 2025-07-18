/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExpectedUeBehaviourData struct {
	CommunicationDurationTime  *int                        `json:"communicationDurationTime,omitempty"`
	PeriodicTime               *int                        `json:"periodicTime,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`
	ScheduledCommunicationType ScheduledCommunicationType  `json:"scheduledCommunicationType,omitempty"`
	TrafficProfile             TrafficProfile              `json:"trafficProfile,omitempty"`
	StationaryIndication       StationaryIndication        `json:"stationaryIndication,omitempty"`
	ExpectedUmts               []LocationArea              `json:"expectedUmts,omitempty"`
	BatteryIndication          *BatteryIndication          `json:"batteryIndication,omitempty"`
	ValidityTime               string                      `json:"validityTime,omitempty"`
}
