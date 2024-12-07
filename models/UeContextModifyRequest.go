/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:22 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextModifyRequest struct {
	Rfsp                 *int64          `json:"rfsp,omitempty"`
	OldAmfNgapId         *int64          `json:"oldAmfNgapId,omitempty"`
	RrcStatusReport      *int16          `json:"rrcStatusReport,omitempty"`
	EmergencyFallbackInd *bool           `json:"emergencyFallbackInd,omitempty"`
	PagingPriority       *int16          `json:"pagingPriority,omitempty"`
	CoreAssistInfo       *CoreAssistInfo `json:"coreAssistInfo,omitempty"`
	UeAmbr               *UeAmbr         `json:"ueAmbr,omitempty"`
	SecKey               []byte          `json:"secKey,omitempty"`
}
