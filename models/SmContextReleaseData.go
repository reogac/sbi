/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Wed Apr 30 14:54:40 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type SmContextReleaseData struct {
	NgApCause         *NgApCause       `json:"ngApCause,omitempty"`
	FiveGMmCauseValue *int             `json:"5gMmCauseValue,omitempty"`
	N2SmInfo          *RefToBinaryData `json:"n2SmInfo,omitempty"`
	N2SmInfoType      N2SmInfoType     `json:"n2SmInfoType,omitempty"`
	IsmfReleaseOnly   *bool            `json:"ismfReleaseOnly,omitempty"`
	Cause             Cause            `json:"cause,omitempty"`
	UeTimeZone        string           `json:"ueTimeZone,omitempty"`
	AddUeLocation     *UserLocation    `json:"addUeLocation,omitempty"`
	VsmfReleaseOnly   *bool            `json:"vsmfReleaseOnly,omitempty"`
	UeLocation        *UserLocation    `json:"ueLocation,omitempty"`
}
