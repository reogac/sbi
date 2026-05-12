/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ExpectedUeBehaviour struct {
	BatteryIndication          *BatteryIndication          `json:"batteryIndication,omitempty"`
	PeriodicTime               *int                        `json:"periodicTime,omitempty"`
	ExpectedUmts               []LocationArea              `json:"expectedUmts,omitempty"`
	TrafficProfile             TrafficProfile              `json:"trafficProfile,omitempty"`
	ValidityTime               string                      `json:"validityTime,omitempty"`
	MtcProviderInformation     string                      `json:"mtcProviderInformation,omitempty"`
	AfInstanceId               string                      `json:"afInstanceId"`
	ReferenceId                int                         `json:"referenceId"`
	StationaryIndication       StationaryIndication        `json:"stationaryIndication,omitempty"`
	CommunicationDurationTime  *int                        `json:"communicationDurationTime,omitempty"`
	ScheduledCommunicationType ScheduledCommunicationType  `json:"scheduledCommunicationType,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`
}
