/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jun 17 13:36:02 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PpMaximumResponseTime struct {
	AfInstanceId           string `json:"afInstanceId"`
	ReferenceId            int    `json:"referenceId"`
	ValidityTime           string `json:"validityTime,omitempty"`
	MtcProviderInformation string `json:"mtcProviderInformation,omitempty"`
	MaximumResponseTime    int    `json:"maximumResponseTime"`
}
