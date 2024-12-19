/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 14:25:56 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GeraLocation struct {
	AgeOfLocationInformation *int            `json:"ageOfLocationInformation,omitempty"`
	GeographicalInformation  string          `json:"geographicalInformation,omitempty"`
	GeodeticInformation      string          `json:"geodeticInformation,omitempty"`
	Cgi                      *CellGlobalId   `json:"cgi,omitempty"`
	Sai                      *ServiceAreaId  `json:"sai,omitempty"`
	VlrNumber                string          `json:"vlrNumber,omitempty"`
	MscNumber                string          `json:"mscNumber,omitempty"`
	UeLocationTimestamp      string          `json:"ueLocationTimestamp,omitempty"`
	LocationNumber           string          `json:"locationNumber,omitempty"`
	Lai                      *LocationAreaId `json:"lai,omitempty"`
	Rai                      *RoutingAreaId  `json:"rai,omitempty"`
}
