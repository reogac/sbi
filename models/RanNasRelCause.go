/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:26 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RanNasRelCause struct {
	FiveGMmCause *int       `json:"5gMmCause,omitempty"`
	FiveGSmCause *int       `json:"5gSmCause,omitempty"`
	EpsCause     string     `json:"epsCause,omitempty"`
	NgApCause    *NgApCause `json:"ngApCause,omitempty"`
}
