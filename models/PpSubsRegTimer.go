/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul 22 12:00:37 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PpSubsRegTimer struct {
	ValidityTime           string `json:"validityTime,omitempty"`
	MtcProviderInformation string `json:"mtcProviderInformation,omitempty"`
	SubsRegTimer           int    `json:"subsRegTimer"`
	AfInstanceId           string `json:"afInstanceId"`
	ReferenceId            int    `json:"referenceId"`
}
