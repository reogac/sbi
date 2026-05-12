/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceAreaRestriction struct {
	MaxNumOfTAs                   *int            `json:"maxNumOfTAs,omitempty"`
	MaxNumOfTAsForNotAllowedAreas *int            `json:"maxNumOfTAsForNotAllowedAreas,omitempty"`
	RestrictionType               RestrictionType `json:"restrictionType,omitempty"`
	Areas                         []Area          `json:"areas,omitempty"`
}
