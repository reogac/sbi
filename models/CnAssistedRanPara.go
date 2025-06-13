/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type CnAssistedRanPara struct {
	BatteryIndication          *BatteryIndication          `json:"batteryIndication,omitempty"`
	StationaryIndication       StationaryIndication        `json:"stationaryIndication,omitempty"`
	CommunicationDurationTime  *int                        `json:"communicationDurationTime,omitempty"`
	PeriodicTime               *int                        `json:"periodicTime,omitempty"`
	ScheduledCommunicationTime *ScheduledCommunicationTime `json:"scheduledCommunicationTime,omitempty"`
	ScheduledCommunicationType ScheduledCommunicationType  `json:"scheduledCommunicationType,omitempty"`
	TrafficProfile             TrafficProfile              `json:"trafficProfile,omitempty"`
}
