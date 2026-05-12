/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProvisionedDataSets struct {
	SmData              *SmSubsData                        `json:"smData,omitempty"`
	SmsMngData          *SmsManagementSubscriptionData     `json:"smsMngData,omitempty"`
	LcsPrivacyData      *LcsPrivacyData                    `json:"lcsPrivacyData,omitempty"`
	SmsSubsData         *SmsSubscriptionData               `json:"smsSubsData,omitempty"`
	TraceData           *TraceData                         `json:"traceData,omitempty"`
	V2xData             *V2xSubscriptionData               `json:"v2xData,omitempty"`
	EeProfileData       *EeProfileData                     `json:"eeProfileData,omitempty"`
	AmData              *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	PpProfileData       *PpProfileData                     `json:"ppProfileData,omitempty"`
	LcsMoData           *LcsMoData                         `json:"lcsMoData,omitempty"`
	LcsBcaData          *LcsBroadcastAssistanceTypesData   `json:"lcsBcaData,omitempty"`
	ProseData           *ProseSubscriptionData             `json:"proseData,omitempty"`
	OdbData             *OdbData                           `json:"odbData,omitempty"`
	MbsSubscriptionData *MbsSubscriptionData               `json:"mbsSubscriptionData,omitempty"`
	SmfSelData          *SmfSelectionSubscriptionData      `json:"smfSelData,omitempty"`
}
