/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type CommunicationCharacteristics struct {
	PpActiveTime          *PpActiveTime          `json:"ppActiveTime,omitempty"`
	PpDlPacketCount       *int                   `json:"ppDlPacketCount,omitempty"`
	PpDlPacketCountExt    *PpDlPacketCountExt    `json:"ppDlPacketCountExt,omitempty"`
	PpMaximumResponseTime *PpMaximumResponseTime `json:"ppMaximumResponseTime,omitempty"`
	PpMaximumLatency      *PpMaximumLatency      `json:"ppMaximumLatency,omitempty"`
	PpSubsRegTimer        *PpSubsRegTimer        `json:"ppSubsRegTimer,omitempty"`
}
