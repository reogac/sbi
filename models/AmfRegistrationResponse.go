/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:08 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfRegistrationResponse struct {
	PlmnList   []SupportedPlmnItem `json:"plmnList,omitempty"`
	AmfPointer int16               `json:"amfPointer"`
}
