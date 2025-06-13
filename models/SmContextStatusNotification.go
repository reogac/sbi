/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextStatusNotification struct {
	NotifyCorrelationIdsForddnFailure []string             `json:"notifyCorrelationIdsForddnFailure,omitempty"`
	NewIntermediateSmfId              string               `json:"newIntermediateSmfId,omitempty"`
	AltAnchorSmfId                    string               `json:"altAnchorSmfId,omitempty"`
	NewSmfId                          string               `json:"newSmfId,omitempty"`
	NewSmfSetId                       string               `json:"newSmfSetId,omitempty"`
	OldSmContextRef                   string               `json:"oldSmContextRef,omitempty"`
	OldPduSessionRef                  string               `json:"oldPduSessionRef,omitempty"`
	InterPlmnApiRoot                  string               `json:"interPlmnApiRoot,omitempty"`
	StatusInfo                        StatusInfo           `json:"statusInfo"`
	SmallDataRateStatus               *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	ApnRateStatus                     *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	AltAnchorSmfUri                   string               `json:"altAnchorSmfUri,omitempty"`
	DdnFailureStatus                  *bool                `json:"ddnFailureStatus,omitempty"`
	OldSmfId                          string               `json:"oldSmfId,omitempty"`
	TargetDnaiInfo                    *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
}
