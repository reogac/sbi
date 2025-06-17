/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ApnRateStatus struct {
	RemainPacketsDl   *int   `json:"remainPacketsDl,omitempty"`
	ValidityTime      string `json:"validityTime,omitempty"`
	RemainExReportsUl *int   `json:"remainExReportsUl,omitempty"`
	RemainExReportsDl *int   `json:"remainExReportsDl,omitempty"`
	RemainPacketsUl   *int   `json:"remainPacketsUl,omitempty"`
}
