/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
