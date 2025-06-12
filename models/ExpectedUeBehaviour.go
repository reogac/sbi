/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExpectedUeBehaviour struct {
	ScheduledCommunicationType ScheduledCommunicationType  `json:"scheduledCommunicationType,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`
	ExpectedUmts               []LocationArea              `json:"expectedUmts,omitempty"`
	BatteryIndication          *BatteryIndication          `json:"batteryIndication,omitempty"`
	ValidityTime               string                      `json:"validityTime,omitempty"`
	TrafficProfile             TrafficProfile              `json:"trafficProfile,omitempty"`
	MtcProviderInformation     string                      `json:"mtcProviderInformation,omitempty"`
	AfInstanceId               string                      `json:"afInstanceId"`
	ReferenceId                int                         `json:"referenceId"`
	StationaryIndication       StationaryIndication        `json:"stationaryIndication,omitempty"`
	CommunicationDurationTime  *int                        `json:"communicationDurationTime,omitempty"`
	PeriodicTime               *int                        `json:"periodicTime,omitempty"`
}
