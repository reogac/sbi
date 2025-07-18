/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmsManagementSubscriptionData struct {
	MtSmsSubscribed     *bool      `json:"mtSmsSubscribed,omitempty"`
	MtSmsBarringAll     *bool      `json:"mtSmsBarringAll,omitempty"`
	SharedSmsMngDataIds []string   `json:"sharedSmsMngDataIds,omitempty"`
	TraceData           *TraceData `json:"traceData,omitempty"`
	MoSmsBarringRoaming *bool      `json:"moSmsBarringRoaming,omitempty"`
	SupportedFeatures   string     `json:"supportedFeatures,omitempty"`
	MtSmsBarringRoaming *bool      `json:"mtSmsBarringRoaming,omitempty"`
	MoSmsSubscribed     *bool      `json:"moSmsSubscribed,omitempty"`
	MoSmsBarringAll     *bool      `json:"moSmsBarringAll,omitempty"`
}
