/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SharedData struct {
	SharedDnnConfigurations map[string]DnnConfiguration        `json:"sharedDnnConfigurations,omitempty"`
	SharedTraceData         *TraceData                         `json:"sharedTraceData,omitempty"`
	SharedVnGroupDatas      map[string]VnGroupData             `json:"sharedVnGroupDatas,omitempty"`
	TreatmentInstructions   map[string]string                  `json:"treatmentInstructions,omitempty"`
	SharedSmSubsData        *SessionManagementSubscriptionData `json:"sharedSmSubsData,omitempty"`
	SharedDataId            string                             `json:"sharedDataId"`
	SharedAmData            *AccessAndMobilitySubscriptionData `json:"sharedAmData,omitempty"`
	SharedSmsSubsData       *SmsSubscriptionData               `json:"sharedSmsSubsData,omitempty"`
	SharedSmsMngSubsData    *SmsManagementSubscriptionData     `json:"sharedSmsMngSubsData,omitempty"`
	SharedSnssaiInfos       map[string]SnssaiInfo              `json:"sharedSnssaiInfos,omitempty"`
	SharedEcsAddrConfigInfo *EcsAddrConfigInfo                 `json:"sharedEcsAddrConfigInfo,omitempty"`
}
