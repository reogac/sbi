/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:44:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GeraLocation struct {
	LocationNumber           string          `json:"locationNumber,omitempty"`
	Rai                      *RoutingAreaId  `json:"rai,omitempty"`
	MscNumber                string          `json:"mscNumber,omitempty"`
	GeographicalInformation  string          `json:"geographicalInformation,omitempty"`
	Cgi                      *CellGlobalId   `json:"cgi,omitempty"`
	Sai                      *ServiceAreaId  `json:"sai,omitempty"`
	Lai                      *LocationAreaId `json:"lai,omitempty"`
	VlrNumber                string          `json:"vlrNumber,omitempty"`
	AgeOfLocationInformation *int            `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      string          `json:"ueLocationTimestamp,omitempty"`
	GeodeticInformation      string          `json:"geodeticInformation,omitempty"`
}
