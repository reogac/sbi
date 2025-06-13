/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:31 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PolicyAssociation struct {
	N2Pc5Pol      *N2InfoContent            `json:"n2Pc5Pol,omitempty"`
	N2Pc5ProSePol *N2InfoContent            `json:"n2Pc5ProSePol,omitempty"`
	Triggers      []string                  `json:"triggers,omitempty"`
	Pras          map[string]PresenceInfo   `json:"pras,omitempty"`
	SuppFeat      string                    `json:"suppFeat"`
	Request       *PolicyAssociationRequest `json:"request,omitempty"`
	UePolicy      string                    `json:"uePolicy,omitempty"`
}
