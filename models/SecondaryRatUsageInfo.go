/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed May 14 15:26:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SecondaryRatUsageInfo struct {
	SecondaryRatType    RatType              `json:"secondaryRatType"`
	QosFlowsUsageData   []QosFlowUsageReport `json:"qosFlowsUsageData,omitempty"`
	PduSessionUsageData []VolumeTimedReport  `json:"pduSessionUsageData,omitempty"`
}
