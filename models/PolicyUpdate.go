/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:38 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyUpdate struct {
	Triggers      []string                `json:"triggers,omitempty"`
	Pras          map[string]PresenceInfo `json:"pras,omitempty"`
	ResourceUri   string                  `json:"resourceUri"`
	UePolicy      string                  `json:"uePolicy,omitempty"`
	N2Pc5Pol      *N2InfoContent          `json:"n2Pc5Pol,omitempty"`
	N2Pc5ProSePol *N2InfoContent          `json:"n2Pc5ProSePol,omitempty"`
}
