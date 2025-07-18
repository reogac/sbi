/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SequenceNumber struct {
	DifSign     Sign           `json:"difSign,omitempty"`
	SqnScheme   SqnScheme      `json:"sqnScheme,omitempty"`
	Sqn         string         `json:"sqn,omitempty"`
	LastIndexes map[string]int `json:"lastIndexes,omitempty"`
	IndLength   *int           `json:"indLength,omitempty"`
}
