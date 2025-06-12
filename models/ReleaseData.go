/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Jun 12 16:32:20 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ReleaseData struct {
	AddUeLocation           *UserLocation             `json:"addUeLocation,omitempty"`
	SecondaryRatUsageInfo   []SecondaryRatUsageInfo   `json:"secondaryRatUsageInfo,omitempty"`
	N4Info                  *N4Information            `json:"n4Info,omitempty"`
	N4InfoExt1              *N4Information            `json:"n4InfoExt1,omitempty"`
	N4InfoExt2              *N4Information            `json:"n4InfoExt2,omitempty"`
	Cause                   Cause                     `json:"cause,omitempty"`
	FiveGMmCauseValue       *int                      `json:"5gMmCauseValue,omitempty"`
	UeTimeZone              string                    `json:"ueTimeZone,omitempty"`
	SecondaryRatUsageReport []SecondaryRatUsageReport `json:"secondaryRatUsageReport,omitempty"`
	NgApCause               *NgApCause                `json:"ngApCause,omitempty"`
	UeLocation              *UserLocation             `json:"ueLocation,omitempty"`
}
