/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:44 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDeleteData struct {
	AccuUsageReports     []AccuUsageReport  `json:"accuUsageReports,omitempty"`
	PduSessRelCause      PduSessionRelCause `json:"pduSessRelCause,omitempty"`
	UserLocationInfo     *UserLocation      `json:"userLocationInfo,omitempty"`
	UeTimeZone           string             `json:"ueTimeZone,omitempty"`
	ServingNetwork       *PlmnIdNid         `json:"servingNetwork,omitempty"`
	UserLocationInfoTime string             `json:"userLocationInfoTime,omitempty"`
	RanNasRelCauses      []RanNasRelCause   `json:"ranNasRelCauses,omitempty"`
}
