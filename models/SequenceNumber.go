/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SequenceNumber struct {
	SqnScheme   SqnScheme      `json:"sqnScheme,omitempty"`
	Sqn         string         `json:"sqn,omitempty"`
	LastIndexes map[string]int `json:"lastIndexes,omitempty"`
	IndLength   *int           `json:"indLength,omitempty"`
	DifSign     Sign           `json:"difSign,omitempty"`
}
