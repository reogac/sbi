/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ServiceAreaRestriction struct {
	MaxNumOfTAs                   *int            `json:"maxNumOfTAs,omitempty"`
	MaxNumOfTAsForNotAllowedAreas *int            `json:"maxNumOfTAsForNotAllowedAreas,omitempty"`
	RestrictionType               RestrictionType `json:"restrictionType,omitempty"`
	Areas                         []Area          `json:"areas,omitempty"`
}
