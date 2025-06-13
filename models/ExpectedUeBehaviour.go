/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExpectedUeBehaviour struct {
	StationaryIndication       StationaryIndication        `json:"stationaryIndication,omitempty"`
	ScheduledCommunicationType ScheduledCommunicationType  `json:"scheduledCommunicationType,omitempty"`
	PeriodicTime               *int                        `json:"periodicTime,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`
	ValidityTime               string                      `json:"validityTime,omitempty"`
	AfInstanceId               string                      `json:"afInstanceId"`
	CommunicationDurationTime  *int                        `json:"communicationDurationTime,omitempty"`
	ExpectedUmts               []LocationArea              `json:"expectedUmts,omitempty"`
	TrafficProfile             TrafficProfile              `json:"trafficProfile,omitempty"`
	BatteryIndication          *BatteryIndication          `json:"batteryIndication,omitempty"`
	MtcProviderInformation     string                      `json:"mtcProviderInformation,omitempty"`
	ReferenceId                int                         `json:"referenceId"`
}
