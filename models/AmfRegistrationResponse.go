/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Aug 26 11:15:54 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmfRegistrationResponse struct {
	Lease      *Lease `json:"lease,omitempty"`
	AmfPointer int16  `json:"amfPointer"`
}
