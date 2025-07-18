/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NnwdafEventsSubscription struct {
	EventSubscriptions []EventSubscription    `json:"eventSubscriptions"`
	EventNotifications []EventNotification    `json:"eventNotifications,omitempty"`
	PrevSub            *PrevSubInfo           `json:"prevSub,omitempty"`
	EvtReq             *ReportingInformation  `json:"evtReq,omitempty"`
	NotificationURI    string                 `json:"notificationURI,omitempty"`
	NotifCorrId        string                 `json:"notifCorrId,omitempty"`
	SupportedFeatures  string                 `json:"supportedFeatures,omitempty"`
	FailEventReports   []FailureEventInfo     `json:"failEventReports,omitempty"`
	ConsNfInfo         *ConsumerNfInformation `json:"consNfInfo,omitempty"`
}
