/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextInSmfData struct {
	PgwInfo       []PgwInfo             `json:"pgwInfo,omitempty"`
	EmergencyInfo *EmergencyInfo        `json:"emergencyInfo,omitempty"`
	PduSessions   map[string]PduSession `json:"pduSessions,omitempty"`
}
