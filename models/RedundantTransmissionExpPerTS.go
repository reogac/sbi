/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:10 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RedundantTransmissionExpPerTS struct {
	RedTransStatus  *bool                     `json:"redTransStatus,omitempty"`
	UeRatio         *int                      `json:"ueRatio,omitempty"`
	Confidence      *int                      `json:"confidence,omitempty"`
	TsStart         string                    `json:"tsStart"`
	TsDuration      int                       `json:"tsDuration"`
	ObsvRedTransExp ObservedRedundantTransExp `json:"obsvRedTransExp"`
}
