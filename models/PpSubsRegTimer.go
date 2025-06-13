/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PpSubsRegTimer struct {
	SubsRegTimer           int    `json:"subsRegTimer"`
	AfInstanceId           string `json:"afInstanceId"`
	ReferenceId            int    `json:"referenceId"`
	ValidityTime           string `json:"validityTime,omitempty"`
	MtcProviderInformation string `json:"mtcProviderInformation,omitempty"`
}
