/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GeraLocation struct {
	MscNumber                string          `json:"mscNumber,omitempty"`
	Sai                      *ServiceAreaId  `json:"sai,omitempty"`
	Cgi                      *CellGlobalId   `json:"cgi,omitempty"`
	Lai                      *LocationAreaId `json:"lai,omitempty"`
	Rai                      *RoutingAreaId  `json:"rai,omitempty"`
	VlrNumber                string          `json:"vlrNumber,omitempty"`
	AgeOfLocationInformation *int            `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      string          `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation  string          `json:"geographicalInformation,omitempty"`
	LocationNumber           string          `json:"locationNumber,omitempty"`
	GeodeticInformation      string          `json:"geodeticInformation,omitempty"`
}
