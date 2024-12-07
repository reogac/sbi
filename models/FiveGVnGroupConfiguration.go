/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type FiveGVnGroupConfiguration struct {
	InternalGroupIdentifier string            `json:"internalGroupIdentifier,omitempty"`
	MtcProviderInformation  string            `json:"mtcProviderInformation,omitempty"`
	FiveGVnGroupData        *FiveGVnGroupData `json:"5gVnGroupData,omitempty"`
	Members                 []string          `json:"members,omitempty"`
	ReferenceId             *int              `json:"referenceId,omitempty"`
	AfInstanceId            string            `json:"afInstanceId,omitempty"`
}
