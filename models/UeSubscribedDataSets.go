/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeSubscribedDataSets struct {
	V2xData             *V2xSubscriptionData               `json:"v2xData,omitempty"`
	ProseData           *ProseSubscriptionData             `json:"proseData,omitempty"`
	MbsSubscriptionData *MbsSubscriptionData               `json:"mbsSubscriptionData,omitempty"`
	SmfSelData          *SmfSelectionSubscriptionData      `json:"smfSelData,omitempty"`
	LcsPrivacyData      *LcsPrivacyData                    `json:"lcsPrivacyData,omitempty"`
	LcsBcaData          *LcsBroadcastAssistanceTypesData   `json:"lcsBcaData,omitempty"`
	AmData              *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	SmData              *SmSubsData                        `json:"smData,omitempty"`
	TraceData           *TraceData                         `json:"traceData,omitempty"`
	SmsMngData          *SmsManagementSubscriptionData     `json:"smsMngData,omitempty"`
	LcsMoData           *LcsMoData                         `json:"lcsMoData,omitempty"`
	OdbData             *OdbData                           `json:"odbData,omitempty"`
	SmsSubsData         *SmsSubscriptionData               `json:"smsSubsData,omitempty"`
	EeProfileData       *EeProfileData                     `json:"eeProfileData,omitempty"`
	PpProfileData       *PpProfileData                     `json:"ppProfileData,omitempty"`
	UcData              *UcSubscriptionData                `json:"ucData,omitempty"`
}
