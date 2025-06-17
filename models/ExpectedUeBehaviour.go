/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExpectedUeBehaviour struct {
	StationaryIndication       StationaryIndication        `json:"stationaryIndication,omitempty"`
	ScheduledCommunicationType ScheduledCommunicationType  `json:"scheduledCommunicationType,omitempty"`
	ExpectedUmts               []LocationArea              `json:"expectedUmts,omitempty"`
	BatteryIndication          *BatteryIndication          `json:"batteryIndication,omitempty"`
	ValidityTime               string                      `json:"validityTime,omitempty"`
	MtcProviderInformation     string                      `json:"mtcProviderInformation,omitempty"`
	AfInstanceId               string                      `json:"afInstanceId"`
	ReferenceId                int                         `json:"referenceId"`
	CommunicationDurationTime  *int                        `json:"communicationDurationTime,omitempty"`
	PeriodicTime               *int                        `json:"periodicTime,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`
	TrafficProfile             TrafficProfile              `json:"trafficProfile,omitempty"`
}
