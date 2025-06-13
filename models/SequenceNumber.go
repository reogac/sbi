/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
