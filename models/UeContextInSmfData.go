/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:39 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UeContextInSmfData struct {
	PduSessions   map[string]PduSession `json:"pduSessions,omitempty"`
	PgwInfo       []PgwInfo             `json:"pgwInfo,omitempty"`
	EmergencyInfo *EmergencyInfo        `json:"emergencyInfo,omitempty"`
}
