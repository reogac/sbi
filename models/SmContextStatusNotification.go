/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Apr 30 14:54:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextStatusNotification struct {
	AltAnchorSmfId                    string               `json:"altAnchorSmfId,omitempty"`
	SmallDataRateStatus               *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	DdnFailureStatus                  *bool                `json:"ddnFailureStatus,omitempty"`
	NotifyCorrelationIdsForddnFailure []string             `json:"notifyCorrelationIdsForddnFailure,omitempty"`
	OldSmfId                          string               `json:"oldSmfId,omitempty"`
	StatusInfo                        StatusInfo           `json:"statusInfo"`
	ApnRateStatus                     *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	NewIntermediateSmfId              string               `json:"newIntermediateSmfId,omitempty"`
	AltAnchorSmfUri                   string               `json:"altAnchorSmfUri,omitempty"`
	TargetDnaiInfo                    *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
	OldPduSessionRef                  string               `json:"oldPduSessionRef,omitempty"`
	NewSmfId                          string               `json:"newSmfId,omitempty"`
	NewSmfSetId                       string               `json:"newSmfSetId,omitempty"`
	OldSmContextRef                   string               `json:"oldSmContextRef,omitempty"`
	InterPlmnApiRoot                  string               `json:"interPlmnApiRoot,omitempty"`
}
