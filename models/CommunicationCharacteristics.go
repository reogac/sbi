/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type CommunicationCharacteristics struct {
	PpSubsRegTimer        *PpSubsRegTimer        `json:"ppSubsRegTimer,omitempty"`
	PpActiveTime          *PpActiveTime          `json:"ppActiveTime,omitempty"`
	PpDlPacketCount       *int                   `json:"ppDlPacketCount,omitempty"`
	PpDlPacketCountExt    *PpDlPacketCountExt    `json:"ppDlPacketCountExt,omitempty"`
	PpMaximumResponseTime *PpMaximumResponseTime `json:"ppMaximumResponseTime,omitempty"`
	PpMaximumLatency      *PpMaximumLatency      `json:"ppMaximumLatency,omitempty"`
}
