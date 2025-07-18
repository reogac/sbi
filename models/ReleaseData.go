/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReleaseData struct {
	FiveGMmCauseValue       *int                      `json:"5gMmCauseValue,omitempty"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	AddUeLocation           *UserLocation             `json:"addUeLocation,omitempty"`
	SecondaryRatUsageReport []SecondaryRatUsageReport `json:"secondaryRatUsageReport,omitempty"`
	N4InfoExt2              *N4Information            `json:"n4InfoExt2,omitempty"`
	Cause                   Cause                     `json:"cause,omitempty"`
	NgApCause               *NgApCause                `json:"ngApCause,omitempty"`
	UeLocation              *UserLocation             `json:"ueLocation,omitempty"`
	SecondaryRatUsageInfo   []SecondaryRatUsageInfo   `json:"secondaryRatUsageInfo,omitempty"`
	N4Info                  *N4Information            `json:"n4Info,omitempty"`
	N4InfoExt1              *N4Information            `json:"n4InfoExt1,omitempty"`
}
