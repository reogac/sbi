/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ApnRateStatus struct {
	RemainPacketsUl   *int   `json:"remainPacketsUl,omitempty"`
	RemainPacketsDl   *int   `json:"remainPacketsDl,omitempty"`
	ValidityTime      string `json:"validityTime,omitempty"`
	RemainExReportsUl *int   `json:"remainExReportsUl,omitempty"`
	RemainExReportsDl *int   `json:"remainExReportsDl,omitempty"`
}
