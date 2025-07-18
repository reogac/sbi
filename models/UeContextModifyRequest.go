/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextModifyRequest struct {
	EmergencyFallbackInd *bool           `json:"emergencyFallbackInd,omitempty"`
	PagingPriority       *int16          `json:"pagingPriority,omitempty"`
	CoreAssistInfo       *CoreAssistInfo `json:"coreAssistInfo,omitempty"`
	UeAmbr               *UeAmbr         `json:"ueAmbr,omitempty"`
	SecKey               []byte          `json:"secKey,omitempty"`
	Rfsp                 *int64          `json:"rfsp,omitempty"`
	OldAmfNgapId         *int64          `json:"oldAmfNgapId,omitempty"`
	RrcStatusReport      *int16          `json:"rrcStatusReport,omitempty"`
}
