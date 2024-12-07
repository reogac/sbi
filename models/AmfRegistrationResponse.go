/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:13 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfRegistrationResponse struct {
	AmfPointer int16    `json:"amfPointer"`
	Slices     []Snssai `json:"slices,omitempty"`
}
