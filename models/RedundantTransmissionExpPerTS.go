/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:26 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RedundantTransmissionExpPerTS struct {
	TsStart         string                    `json:"tsStart"`
	TsDuration      int                       `json:"tsDuration"`
	ObsvRedTransExp ObservedRedundantTransExp `json:"obsvRedTransExp"`
	RedTransStatus  *bool                     `json:"redTransStatus,omitempty"`
	UeRatio         *int                      `json:"ueRatio,omitempty"`
	Confidence      *int                      `json:"confidence,omitempty"`
}
