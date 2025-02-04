/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Feb  4 19:25:07 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfRegistrationResponse struct {
	Slices     []Snssai `json:"slices,omitempty"`
	AmfPointer int16    `json:"amfPointer"`
}
