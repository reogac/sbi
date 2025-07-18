/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:35 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextStatusNotification struct {
	ApnRateStatus                     *ApnRateStatus       `json:"apnRateStatus,omitempty"`
	NewSmfSetId                       string               `json:"newSmfSetId,omitempty"`
	OldPduSessionRef                  string               `json:"oldPduSessionRef,omitempty"`
	InterPlmnApiRoot                  string               `json:"interPlmnApiRoot,omitempty"`
	NewSmfId                          string               `json:"newSmfId,omitempty"`
	OldSmfId                          string               `json:"oldSmfId,omitempty"`
	OldSmContextRef                   string               `json:"oldSmContextRef,omitempty"`
	NewIntermediateSmfId              string               `json:"newIntermediateSmfId,omitempty"`
	AltAnchorSmfId                    string               `json:"altAnchorSmfId,omitempty"`
	StatusInfo                        StatusInfo           `json:"statusInfo"`
	SmallDataRateStatus               *SmallDataRateStatus `json:"smallDataRateStatus,omitempty"`
	DdnFailureStatus                  *bool                `json:"ddnFailureStatus,omitempty"`
	NotifyCorrelationIdsForddnFailure []string             `json:"notifyCorrelationIdsForddnFailure,omitempty"`
	AltAnchorSmfUri                   string               `json:"altAnchorSmfUri,omitempty"`
	TargetDnaiInfo                    *TargetDnaiInfo      `json:"targetDnaiInfo,omitempty"`
}
