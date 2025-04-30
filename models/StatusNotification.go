/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Apr 30 14:54:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type StatusNotification struct {
	TargetDnaiInfo      *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
	OldPduSessionRef    string               `json:"oldPduSessionRef,omitempty"`
	InterPlmnApiRoot    string               `json:"interPlmnApiRoot,omitempty"`
	StatusInfo          StatusInfo           `json:"statusInfo"`
	ApnRateStatus       *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	NewSmfId            string               `json:"newSmfId,omitempty"`
	EpsPdnCnxInfo       *EpsPdnCnxInfo       `json:"epsPdnCnxInfo,omitempty"`
	IntraPlmnApiRoot    string               `json:"intraPlmnApiRoot,omitempty"`
	SmallDataRateStatus *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
}
