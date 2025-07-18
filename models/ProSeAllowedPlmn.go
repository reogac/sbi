/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:49 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProSeAllowedPlmn struct {
	VisitedPlmn        PlmnId   `json:"visitedPlmn"`
	ProseDirectAllowed []string `json:"proseDirectAllowed,omitempty"`
}
