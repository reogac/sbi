/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:45 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PpDlPacketCountExt struct {
	MtcProviderInformation string  `json:"mtcProviderInformation,omitempty"`
	AfInstanceId           string  `json:"afInstanceId"`
	ReferenceId            int     `json:"referenceId"`
	Dnn                    string  `json:"dnn,omitempty"`
	SingleNssai            *Snssai `json:"singleNssai,omitempty"`
	ValidityTime           string  `json:"validityTime,omitempty"`
}
