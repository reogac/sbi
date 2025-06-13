/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:28 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NrLocation struct {
	Tai                      Tai              `json:"tai"`
	Ncgi                     Ncgi             `json:"ncgi"`
	IgnoreNcgi               *bool            `json:"ignoreNcgi,omitempty"`
	AgeOfLocationInformation *int             `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      string           `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation  string           `json:"geographicalInformation,omitempty"`
	GeodeticInformation      string           `json:"geodeticInformation,omitempty"`
	GlobalGnbId              *GlobalRanNodeId `json:"globalGnbId,omitempty"`
}
