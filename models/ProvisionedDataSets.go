/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProvisionedDataSets struct {
	EeProfileData       *EeProfileData                     `json:"eeProfileData,omitempty"`
	SmData              *SmSubsData                        `json:"smData,omitempty"`
	TraceData           *TraceData                         `json:"traceData,omitempty"`
	SmsMngData          *SmsManagementSubscriptionData     `json:"smsMngData,omitempty"`
	LcsMoData           *LcsMoData                         `json:"lcsMoData,omitempty"`
	OdbData             *OdbData                           `json:"odbData,omitempty"`
	AmData              *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	SmfSelData          *SmfSelectionSubscriptionData      `json:"smfSelData,omitempty"`
	LcsPrivacyData      *LcsPrivacyData                    `json:"lcsPrivacyData,omitempty"`
	V2xData             *V2xSubscriptionData               `json:"v2xData,omitempty"`
	PpProfileData       *PpProfileData                     `json:"ppProfileData,omitempty"`
	SmsSubsData         *SmsSubscriptionData               `json:"smsSubsData,omitempty"`
	LcsBcaData          *LcsBroadcastAssistanceTypesData   `json:"lcsBcaData,omitempty"`
	ProseData           *ProseSubscriptionData             `json:"proseData,omitempty"`
	MbsSubscriptionData *MbsSubscriptionData               `json:"mbsSubscriptionData,omitempty"`
}
