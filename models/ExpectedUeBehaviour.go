/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExpectedUeBehaviour struct {
	ReferenceId                int                         `json:"referenceId"`
	StationaryIndication       StationaryIndication        `json:"stationaryIndication,omitempty"`
	ExpectedUmts               []LocationArea              `json:"expectedUmts,omitempty"`
	TrafficProfile             TrafficProfile              `json:"trafficProfile,omitempty"`
	BatteryIndication          *BatteryIndication          `json:"batteryIndication,omitempty"`
	MtcProviderInformation     string                      `json:"mtcProviderInformation,omitempty"`
	AfInstanceId               string                      `json:"afInstanceId"`
	ScheduledCommunicationType ScheduledCommunicationType  `json:"scheduledCommunicationType,omitempty"`
	PeriodicTime               *int                        `json:"periodicTime,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`
	ValidityTime               string                      `json:"validityTime,omitempty"`
	CommunicationDurationTime  *int                        `json:"communicationDurationTime,omitempty"`
}
