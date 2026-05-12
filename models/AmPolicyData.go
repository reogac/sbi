/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:46 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmPolicyData struct {
	PraInfos  map[string]PresenceInfo `json:"praInfos,omitempty"`
	SubscCats []string                `json:"subscCats,omitempty"`
}
