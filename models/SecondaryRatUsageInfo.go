/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SecondaryRatUsageInfo struct {
	QosFlowsUsageData   []QosFlowUsageReport `json:"qosFlowsUsageData,omitempty"`
	PduSessionUsageData []VolumeTimedReport  `json:"pduSessionUsageData,omitempty"`
	SecondaryRatType    RatType              `json:"secondaryRatType"`
}
