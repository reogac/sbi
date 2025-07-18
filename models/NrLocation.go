/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NrLocation struct {
	Ncgi                     Ncgi             `json:"ncgi"`
	IgnoreNcgi               *bool            `json:"ignoreNcgi,omitempty"`
	AgeOfLocationInformation *int             `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      string           `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation  string           `json:"geographicalInformation,omitempty"`
	GeodeticInformation      string           `json:"geodeticInformation,omitempty"`
	GlobalGnbId              *GlobalRanNodeId `json:"globalGnbId,omitempty"`
	Tai                      Tai              `json:"tai"`
}
