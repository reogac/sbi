/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type CommunicationCharacteristicsAF struct {
	PpDlPacketCount     *int `json:"ppDlPacketCount,omitempty"`
	MaximumResponseTime *int `json:"maximumResponseTime,omitempty"`
	MaximumLatency      *int `json:"maximumLatency,omitempty"`
}
