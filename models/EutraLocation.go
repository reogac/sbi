/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:28 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EutraLocation struct {
	Tai                      Tai              `json:"tai"`
	AgeOfLocationInformation *int             `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      string           `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation  string           `json:"geographicalInformation,omitempty"`
	GeodeticInformation      string           `json:"geodeticInformation,omitempty"`
	IgnoreTai                *bool            `json:"ignoreTai,omitempty"`
	Ecgi                     Ecgi             `json:"ecgi"`
	IgnoreEcgi               *bool            `json:"ignoreEcgi,omitempty"`
	GlobalNgenbId            *GlobalRanNodeId `json:"globalNgenbId,omitempty"`
	GlobalENbId              *GlobalRanNodeId `json:"globalENbId,omitempty"`
}
