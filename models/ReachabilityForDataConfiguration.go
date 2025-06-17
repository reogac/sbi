/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReachabilityForDataConfiguration struct {
	ReportCfg   ReachabilityForDataReportConfig `json:"reportCfg"`
	MinInterval *int                            `json:"minInterval,omitempty"`
}
