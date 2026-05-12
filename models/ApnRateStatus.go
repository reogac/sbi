/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:30 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ApnRateStatus struct {
	ValidityTime      string `json:"validityTime,omitempty"`
	RemainExReportsUl *int   `json:"remainExReportsUl,omitempty"`
	RemainExReportsDl *int   `json:"remainExReportsDl,omitempty"`
	RemainPacketsUl   *int   `json:"remainPacketsUl,omitempty"`
	RemainPacketsDl   *int   `json:"remainPacketsDl,omitempty"`
}
