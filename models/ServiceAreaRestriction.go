/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceAreaRestriction struct {
	RestrictionType               RestrictionType `json:"restrictionType,omitempty"`
	Areas                         []Area          `json:"areas,omitempty"`
	MaxNumOfTAs                   *int            `json:"maxNumOfTAs,omitempty"`
	MaxNumOfTAsForNotAllowedAreas *int            `json:"maxNumOfTAsForNotAllowedAreas,omitempty"`
}
