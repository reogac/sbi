/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:41 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeCommunication struct {
	RecurringTime     *ScheduledCommunicationTime `json:"recurringTime,omitempty"`
	Ratio             *int                        `json:"ratio,omitempty"`
	Confidence        *int                        `json:"confidence,omitempty"`
	PerioTime         *int                        `json:"perioTime,omitempty"`
	PerioTimeVariance *float64                    `json:"perioTimeVariance,omitempty"`
	Ts                string                      `json:"ts,omitempty"`
	TrafChar          *TrafficCharacterization    `json:"trafChar,omitempty"`
	PerioCommInd      *bool                       `json:"perioCommInd,omitempty"`
	AnaOfAppList      *AppListForUeComm           `json:"anaOfAppList,omitempty"`
	SessInactTimer    *SessInactTimerForUeComm    `json:"sessInactTimer,omitempty"`
	CommDur           *int                        `json:"commDur,omitempty"`
	CommDurVariance   *float64                    `json:"commDurVariance,omitempty"`
	TsVariance        *float64                    `json:"tsVariance,omitempty"`
}
