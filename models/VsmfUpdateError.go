/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:30 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type VsmfUpdateError struct {
	Pti                   *int              `json:"pti,omitempty"`
	N1smCause             string            `json:"n1smCause,omitempty"`
	NgApCause             *NgApCause        `json:"ngApCause,omitempty"`
	FiveGMmCauseValue     *int              `json:"5gMmCauseValue,omitempty"`
	RecoveryTime          string            `json:"recoveryTime,omitempty"`
	N4Info                *N4Information    `json:"n4Info,omitempty"`
	N4InfoExt3            *N4Information    `json:"n4InfoExt3,omitempty"`
	Error                 ExtProblemDetails `json:"error"`
	N1SmInfoFromUe        *RefToBinaryData  `json:"n1SmInfoFromUe,omitempty"`
	UnknownN1SmInfo       *RefToBinaryData  `json:"unknownN1SmInfo,omitempty"`
	FailedToAssignEbiList []Arp             `json:"failedToAssignEbiList,omitempty"`
	N4InfoExt1            *N4Information    `json:"n4InfoExt1,omitempty"`
	N4InfoExt2            *N4Information    `json:"n4InfoExt2,omitempty"`
}
