/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Dec 19 15:44:19 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type EutraLocation struct {
	GlobalENbId              *GlobalRanNodeId `json:"globalENbId,omitempty"`
	Tai                      Tai              `json:"tai"`
	IgnoreTai                *bool            `json:"ignoreTai,omitempty"`
	GlobalNgenbId            *GlobalRanNodeId `json:"globalNgenbId,omitempty"`
	UeLocationTimestamp      string           `json:"ueLocationTimestamp,omitempty"`
	GeographicalInformation  string           `json:"geographicalInformation,omitempty"`
	GeodeticInformation      string           `json:"geodeticInformation,omitempty"`
	Ecgi                     Ecgi             `json:"ecgi"`
	IgnoreEcgi               *bool            `json:"ignoreEcgi,omitempty"`
	AgeOfLocationInformation *int             `json:"ageOfLocationInformation,omitempty"`
}
