/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
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
