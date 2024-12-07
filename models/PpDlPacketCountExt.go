/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:33 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PpDlPacketCountExt struct {
	SingleNssai            *Snssai `json:"singleNssai,omitempty"`
	ValidityTime           string  `json:"validityTime,omitempty"`
	MtcProviderInformation string  `json:"mtcProviderInformation,omitempty"`
	AfInstanceId           string  `json:"afInstanceId"`
	ReferenceId            int     `json:"referenceId"`
	Dnn                    string  `json:"dnn,omitempty"`
}
