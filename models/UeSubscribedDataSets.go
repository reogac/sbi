/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeSubscribedDataSets struct {
	MbsSubscriptionData *MbsSubscriptionData               `json:"mbsSubscriptionData,omitempty"`
	AmData              *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	V2xData             *V2xSubscriptionData               `json:"v2xData,omitempty"`
	LcsPrivacyData      *LcsPrivacyData                    `json:"lcsPrivacyData,omitempty"`
	ProseData           *ProseSubscriptionData             `json:"proseData,omitempty"`
	OdbData             *OdbData                           `json:"odbData,omitempty"`
	PpProfileData       *PpProfileData                     `json:"ppProfileData,omitempty"`
	SmfSelData          *SmfSelectionSubscriptionData      `json:"smfSelData,omitempty"`
	TraceData           *TraceData                         `json:"traceData,omitempty"`
	SmsMngData          *SmsManagementSubscriptionData     `json:"smsMngData,omitempty"`
	LcsMoData           *LcsMoData                         `json:"lcsMoData,omitempty"`
	EeProfileData       *EeProfileData                     `json:"eeProfileData,omitempty"`
	SmsSubsData         *SmsSubscriptionData               `json:"smsSubsData,omitempty"`
	SmData              *SmSubsData                        `json:"smData,omitempty"`
	UcData              *UcSubscriptionData                `json:"ucData,omitempty"`
	LcsBcaData          *LcsBroadcastAssistanceTypesData   `json:"lcsBcaData,omitempty"`
}
