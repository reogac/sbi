/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:35:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfRegistrationResponse struct {
	AmfPointer int16               `json:"amfPointer"`
	PlmnList   []SupportedPlmnItem `json:"plmnList,omitempty"`
}
