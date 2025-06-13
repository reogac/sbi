/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type FiveGVnGroupConfiguration struct {
	MtcProviderInformation  string            `json:"mtcProviderInformation,omitempty"`
	FiveGVnGroupData        *FiveGVnGroupData `json:"5gVnGroupData,omitempty"`
	Members                 []string          `json:"members,omitempty"`
	ReferenceId             *int              `json:"referenceId,omitempty"`
	AfInstanceId            string            `json:"afInstanceId,omitempty"`
	InternalGroupIdentifier string            `json:"internalGroupIdentifier,omitempty"`
}
