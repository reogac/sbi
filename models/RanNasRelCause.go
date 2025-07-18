/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RanNasRelCause struct {
	NgApCause    *NgApCause `json:"ngApCause,omitempty"`
	FiveGMmCause *int       `json:"5gMmCause,omitempty"`
	FiveGSmCause *int       `json:"5gSmCause,omitempty"`
	EpsCause     string     `json:"epsCause,omitempty"`
}
