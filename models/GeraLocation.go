/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Feb 18 15:01:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type GeraLocation struct {
	AgeOfLocationInformation *int            `json:"ageOfLocationInformation,omitempty"`
	GeographicalInformation  string          `json:"geographicalInformation,omitempty"`
	GeodeticInformation      string          `json:"geodeticInformation,omitempty"`
	Sai                      *ServiceAreaId  `json:"sai,omitempty"`
	Rai                      *RoutingAreaId  `json:"rai,omitempty"`
	MscNumber                string          `json:"mscNumber,omitempty"`
	VlrNumber                string          `json:"vlrNumber,omitempty"`
	UeLocationTimestamp      string          `json:"ueLocationTimestamp,omitempty"`
	LocationNumber           string          `json:"locationNumber,omitempty"`
	Cgi                      *CellGlobalId   `json:"cgi,omitempty"`
	Lai                      *LocationAreaId `json:"lai,omitempty"`
}
