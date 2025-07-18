/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type VsmfUpdateError struct {
	RecoveryTime          string            `json:"recoveryTime,omitempty"`
	N4Info                *N4Information    `json:"n4Info,omitempty"`
	N4InfoExt1            *N4Information    `json:"n4InfoExt1,omitempty"`
	UnknownN1SmInfo       *RefToBinaryData  `json:"unknownN1SmInfo,omitempty"`
	Pti                   *int              `json:"pti,omitempty"`
	N1smCause             string            `json:"n1smCause,omitempty"`
	N1SmInfoFromUe        *RefToBinaryData  `json:"n1SmInfoFromUe,omitempty"`
	FailedToAssignEbiList []Arp             `json:"failedToAssignEbiList,omitempty"`
	NgApCause             *NgApCause        `json:"ngApCause,omitempty"`
	FiveGMmCauseValue     *int              `json:"5gMmCauseValue,omitempty"`
	N4InfoExt2            *N4Information    `json:"n4InfoExt2,omitempty"`
	Error                 ExtProblemDetails `json:"error"`
	N4InfoExt3            *N4Information    `json:"n4InfoExt3,omitempty"`
}
