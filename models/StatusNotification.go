/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type StatusNotification struct {
	StatusInfo          StatusInfo           `json:"statusInfo"`
	ApnRateStatus       *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	OldPduSessionRef    string               `json:"oldPduSessionRef,omitempty"`
	NewSmfId            string               `json:"newSmfId,omitempty"`
	InterPlmnApiRoot    string               `json:"interPlmnApiRoot,omitempty"`
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	TargetDnaiInfo      *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
	EpsPdnCnxInfo       *EpsPdnCnxInfo       `json:"epsPdnCnxInfo,omitempty"`
	IntraPlmnApiRoot    string               `json:"intraPlmnApiRoot,omitempty"`
}
