/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SubscriptionDataSets struct {
	LcsBroadcastAssistanceTypesData *LcsBroadcastAssistanceTypesData   `json:"lcsBroadcastAssistanceTypesData,omitempty"`
	UecSmsfData                     *UeContextInSmsfData               `json:"uecSmsfData,omitempty"`
	TraceData                       *TraceData                         `json:"traceData,omitempty"`
	SmsMngData                      *SmsManagementSubscriptionData     `json:"smsMngData,omitempty"`
	UcData                          *UcSubscriptionData                `json:"ucData,omitempty"`
	UecSmfData                      *UeContextInSmfData                `json:"uecSmfData,omitempty"`
	SmData                          *SmSubsData                        `json:"smData,omitempty"`
	LcsPrivacyData                  *LcsPrivacyData                    `json:"lcsPrivacyData,omitempty"`
	LcsMoData                       *LcsMoData                         `json:"lcsMoData,omitempty"`
	SmfSelData                      *SmfSelectionSubscriptionData      `json:"smfSelData,omitempty"`
	UecAmfData                      *UeContextInAmfData                `json:"uecAmfData,omitempty"`
	V2xData                         *V2xSubscriptionData               `json:"v2xData,omitempty"`
	ProseData                       *ProseSubscriptionData             `json:"proseData,omitempty"`
	MbsData                         *MbsSubscriptionData               `json:"mbsData,omitempty"`
	AmData                          *AccessAndMobilitySubscriptionData `json:"amData,omitempty"`
	SmsSubsData                     *SmsSubscriptionData               `json:"smsSubsData,omitempty"`
}
