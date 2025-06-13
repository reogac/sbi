/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProSeAllowedPlmn struct {
	ProseDirectAllowed []string `json:"proseDirectAllowed,omitempty"`
	VisitedPlmn        PlmnId   `json:"visitedPlmn"`
}
