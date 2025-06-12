/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:13 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfRegistrationResponse struct {
	AmfPointer int16               `json:"amfPointer"`
	PlmnList   []SupportedPlmnItem `json:"plmnList,omitempty"`
}
