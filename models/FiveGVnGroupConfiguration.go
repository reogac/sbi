/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type FiveGVnGroupConfiguration struct {
	AfInstanceId            string            `json:"afInstanceId,omitempty"`
	InternalGroupIdentifier string            `json:"internalGroupIdentifier,omitempty"`
	MtcProviderInformation  string            `json:"mtcProviderInformation,omitempty"`
	FiveGVnGroupData        *FiveGVnGroupData `json:"5gVnGroupData,omitempty"`
	Members                 []string          `json:"members,omitempty"`
	ReferenceId             *int              `json:"referenceId,omitempty"`
}
