/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:36 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type VsmfUpdateError struct {
	N4InfoExt1            *N4Information    `json:"n4InfoExt1,omitempty"`
	Pti                   *int              `json:"pti,omitempty"`
	N1SmInfoFromUe        *RefToBinaryData  `json:"n1SmInfoFromUe,omitempty"`
	FailedToAssignEbiList []Arp             `json:"failedToAssignEbiList,omitempty"`
	NgApCause             *NgApCause        `json:"ngApCause,omitempty"`
	FiveGMmCauseValue     *int              `json:"5gMmCauseValue,omitempty"`
	RecoveryTime          string            `json:"recoveryTime,omitempty"`
	N4Info                *N4Information    `json:"n4Info,omitempty"`
	Error                 ExtProblemDetails `json:"error"`
	N1smCause             string            `json:"n1smCause,omitempty"`
	UnknownN1SmInfo       *RefToBinaryData  `json:"unknownN1SmInfo,omitempty"`
	N4InfoExt2            *N4Information    `json:"n4InfoExt2,omitempty"`
	N4InfoExt3            *N4Information    `json:"n4InfoExt3,omitempty"`
}
