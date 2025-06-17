/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SequenceNumber struct {
	Sqn         string         `json:"sqn,omitempty"`
	LastIndexes map[string]int `json:"lastIndexes,omitempty"`
	IndLength   *int           `json:"indLength,omitempty"`
	DifSign     Sign           `json:"difSign,omitempty"`
	SqnScheme   SqnScheme      `json:"sqnScheme,omitempty"`
}
