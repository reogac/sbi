/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type LcsPrivacy struct {
	Lpi                    *Lpi   `json:"lpi,omitempty"`
	MtcProviderInformation string `json:"mtcProviderInformation,omitempty"`
	AfInstanceId           string `json:"afInstanceId,omitempty"`
	ReferenceId            *int   `json:"referenceId,omitempty"`
}
