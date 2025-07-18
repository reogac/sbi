/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmPolicyDeleteData struct {
	PduSessRelCause      PduSessionRelCause `json:"pduSessRelCause,omitempty"`
	UserLocationInfo     *UserLocation      `json:"userLocationInfo,omitempty"`
	UeTimeZone           string             `json:"ueTimeZone,omitempty"`
	ServingNetwork       *PlmnIdNid         `json:"servingNetwork,omitempty"`
	UserLocationInfoTime string             `json:"userLocationInfoTime,omitempty"`
	RanNasRelCauses      []RanNasRelCause   `json:"ranNasRelCauses,omitempty"`
	AccuUsageReports     []AccuUsageReport  `json:"accuUsageReports,omitempty"`
}
