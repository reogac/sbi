/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProvisionedDataSets struct {
	AmData              *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	SmfSelData          *SmfSelectionSubscriptionData      `json:"smfSelData,omitempty"`
	SmData              *SmSubsData                        `json:"smData,omitempty"`
	EeProfileData       *EeProfileData                     `json:"eeProfileData,omitempty"`
	SmsSubsData         *SmsSubscriptionData               `json:"smsSubsData,omitempty"`
	TraceData           *TraceData                         `json:"traceData,omitempty"`
	V2xData             *V2xSubscriptionData               `json:"v2xData,omitempty"`
	ProseData           *ProseSubscriptionData             `json:"proseData,omitempty"`
	OdbData             *OdbData                           `json:"odbData,omitempty"`
	LcsPrivacyData      *LcsPrivacyData                    `json:"lcsPrivacyData,omitempty"`
	LcsMoData           *LcsMoData                         `json:"lcsMoData,omitempty"`
	LcsBcaData          *LcsBroadcastAssistanceTypesData   `json:"lcsBcaData,omitempty"`
	SmsMngData          *SmsManagementSubscriptionData     `json:"smsMngData,omitempty"`
	PpProfileData       *PpProfileData                     `json:"ppProfileData,omitempty"`
	MbsSubscriptionData *MbsSubscriptionData               `json:"mbsSubscriptionData,omitempty"`
}
