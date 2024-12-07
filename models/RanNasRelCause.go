/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:30 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RanNasRelCause struct {
	NgApCause    *NgApCause `json:"ngApCause,omitempty"`
	FiveGMmCause *int       `json:"5gMmCause,omitempty"`
	FiveGSmCause *int       `json:"5gSmCause,omitempty"`
	EpsCause     string     `json:"epsCause,omitempty"`
}
