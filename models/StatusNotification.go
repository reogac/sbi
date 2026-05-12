/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:30 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type StatusNotification struct {
	StatusInfo          StatusInfo           `json:"statusInfo"`
	OldPduSessionRef    string               `json:"oldPduSessionRef,omitempty"`
	NewSmfId            string               `json:"newSmfId,omitempty"`
	EpsPdnCnxInfo       *EpsPdnCnxInfo       `json:"epsPdnCnxInfo,omitempty"`
	InterPlmnApiRoot    string               `json:"interPlmnApiRoot,omitempty"`
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	ApnRateStatus       *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	TargetDnaiInfo      *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
	IntraPlmnApiRoot    string               `json:"intraPlmnApiRoot,omitempty"`
}
