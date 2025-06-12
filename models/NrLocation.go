/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:34 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type NrLocation struct {
	IgnoreNcgi               *bool            `json:"ignoreNcgi,omitempty"`
	AgeOfLocationInformation *int             `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      string           `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation  string           `json:"geographicalInformation,omitempty"`
	GeodeticInformation      string           `json:"geodeticInformation,omitempty"`
	GlobalGnbId              *GlobalRanNodeId `json:"globalGnbId,omitempty"`
	Tai                      Tai              `json:"tai"`
	Ncgi                     Ncgi             `json:"ncgi"`
}
