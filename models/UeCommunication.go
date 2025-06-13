/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeCommunication struct {
	Confidence        *int                        `json:"confidence,omitempty"`
	AnaOfAppList      *AppListForUeComm           `json:"anaOfAppList,omitempty"`
	CommDur           *int                        `json:"commDur,omitempty"`
	PerioTime         *int                        `json:"perioTime,omitempty"`
	PerioTimeVariance *float64                    `json:"perioTimeVariance,omitempty"`
	RecurringTime     *ScheduledCommunicationTime `json:"recurringTime,omitempty"`
	Ratio             *int                        `json:"ratio,omitempty"`
	PerioCommInd      *bool                       `json:"perioCommInd,omitempty"`
	SessInactTimer    *SessInactTimerForUeComm    `json:"sessInactTimer,omitempty"`
	CommDurVariance   *float64                    `json:"commDurVariance,omitempty"`
	Ts                string                      `json:"ts,omitempty"`
	TsVariance        *float64                    `json:"tsVariance,omitempty"`
	TrafChar          *TrafficCharacterization    `json:"trafChar,omitempty"`
}
