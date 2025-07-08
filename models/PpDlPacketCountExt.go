/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue Jul  8 13:19:46 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type PpDlPacketCountExt struct {
	Dnn                    string  `json:"dnn,omitempty"`
	SingleNssai            *Snssai `json:"singleNssai,omitempty"`
	ValidityTime           string  `json:"validityTime,omitempty"`
	MtcProviderInformation string  `json:"mtcProviderInformation,omitempty"`
	AfInstanceId           string  `json:"afInstanceId"`
	ReferenceId            int     `json:"referenceId"`
}
