/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 14:25:56 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type UtraLocation struct {
	GeographicalInformation  string          `json:"geographicalInformation,omitempty"`
	GeodeticInformation      string          `json:"geodeticInformation,omitempty"`
	Cgi                      *CellGlobalId   `json:"cgi,omitempty"`
	Sai                      *ServiceAreaId  `json:"sai,omitempty"`
	Lai                      *LocationAreaId `json:"lai,omitempty"`
	Rai                      *RoutingAreaId  `json:"rai,omitempty"`
	AgeOfLocationInformation *int            `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      string          `json:"ueLocationTimestamp,omitempty"`
}
