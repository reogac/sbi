/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:15 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextReleaseData struct {
	VsmfReleaseOnly   *bool            `json:"vsmfReleaseOnly,omitempty"`
	Cause             Cause            `json:"cause,omitempty"`
	UeLocation        *UserLocation    `json:"ueLocation,omitempty"`
	UeTimeZone        string           `json:"ueTimeZone,omitempty"`
	N2SmInfo          *RefToBinaryData `json:"n2SmInfo,omitempty"`
	N2SmInfoType      N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	IsmfReleaseOnly   *bool            `json:"ismfReleaseOnly,omitempty"`
	NgApCause         *NgApCause       `json:"ngApCause,omitempty"`
	FiveGMmCauseValue *int             `json:"5gMmCauseValue,omitempty"`
	AddUeLocation     *UserLocation    `json:"addUeLocation,omitempty"`
}
