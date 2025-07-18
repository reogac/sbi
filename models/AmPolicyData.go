/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:50 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmPolicyData struct {
	PraInfos  map[string]PresenceInfo `json:"praInfos,omitempty"`
	SubscCats []string                `json:"subscCats,omitempty"`
}
