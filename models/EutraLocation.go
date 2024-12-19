/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:49:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EutraLocation struct {
	Tai                      Tai              `json:"tai"`
	Ecgi                     Ecgi             `json:"ecgi"`
	IgnoreEcgi               *bool            `json:"ignoreEcgi,omitempty"`
	GeographicalInformation  string           `json:"geographicalInformation,omitempty"`
	GlobalENbId              *GlobalRanNodeId `json:"globalENbId,omitempty"`
	IgnoreTai                *bool            `json:"ignoreTai,omitempty"`
	AgeOfLocationInformation *int             `json:"ageOfLocationInformation,omitempty"`
	UeLocationTimestamp      string           `json:"ueLocationTimestamp,omitempty"`
	GeodeticInformation      string           `json:"geodeticInformation,omitempty"`
	GlobalNgenbId            *GlobalRanNodeId `json:"globalNgenbId,omitempty"`
}
