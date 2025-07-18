/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type AmRequestedValueRep struct {
	UserLoc           *UserLocation           `json:"userLoc,omitempty"`
	PraStatuses       map[string]PresenceInfo `json:"praStatuses,omitempty"`
	AccessTypes       []string                `json:"accessTypes,omitempty"`
	RatTypes          []string                `json:"ratTypes,omitempty"`
	AllowedSnssais    []Snssai                `json:"allowedSnssais,omitempty"`
	N3gAllowedSnssais []Snssai                `json:"n3gAllowedSnssais,omitempty"`
}
