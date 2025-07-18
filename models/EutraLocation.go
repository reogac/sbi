/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:48 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EutraLocation struct {
	IgnoreEcgi               *bool            `json:"ignoreEcgi,omitempty"`
	UeLocationTimestamp      string           `json:"ueLocationTimestamp,omitempty"`
	GeodeticInformation      string           `json:"geodeticInformation,omitempty"`
	GlobalNgenbId            *GlobalRanNodeId `json:"globalNgenbId,omitempty"`
	IgnoreTai                *bool            `json:"ignoreTai,omitempty"`
	Ecgi                     Ecgi             `json:"ecgi"`
	AgeOfLocationInformation *int             `json:"ageOfLocationInformation,omitempty"`
	GeographicalInformation  string           `json:"geographicalInformation,omitempty"`
	GlobalENbId              *GlobalRanNodeId `json:"globalENbId,omitempty"`
	Tai                      Tai              `json:"tai"`
}
