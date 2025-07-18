/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeSubscribedDataSets struct {
	ProseData           *ProseSubscriptionData             `json:"proseData,omitempty"`
	EeProfileData       *EeProfileData                     `json:"eeProfileData,omitempty"`
	LcsBcaData          *LcsBroadcastAssistanceTypesData   `json:"lcsBcaData,omitempty"`
	V2xData             *V2xSubscriptionData               `json:"v2xData,omitempty"`
	SmsSubsData         *SmsSubscriptionData               `json:"smsSubsData,omitempty"`
	SmData              *SmSubsData                        `json:"smData,omitempty"`
	LcsMoData           *LcsMoData                         `json:"lcsMoData,omitempty"`
	SmsMngData          *SmsManagementSubscriptionData     `json:"smsMngData,omitempty"`
	PpProfileData       *PpProfileData                     `json:"ppProfileData,omitempty"`
	UcData              *UcSubscriptionData                `json:"ucData,omitempty"`
	AmData              *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	SmfSelData          *SmfSelectionSubscriptionData      `json:"smfSelData,omitempty"`
	TraceData           *TraceData                         `json:"traceData,omitempty"`
	LcsPrivacyData      *LcsPrivacyData                    `json:"lcsPrivacyData,omitempty"`
	OdbData             *OdbData                           `json:"odbData,omitempty"`
	MbsSubscriptionData *MbsSubscriptionData               `json:"mbsSubscriptionData,omitempty"`
}
