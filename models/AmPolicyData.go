/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:38 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmPolicyData struct {
	SubscCats []string                `json:"subscCats,omitempty"`
	PraInfos  map[string]PresenceInfo `json:"praInfos,omitempty"`
}
