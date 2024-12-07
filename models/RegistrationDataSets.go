/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Sat Dec  7 16:57:27 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type RegistrationDataSets struct {
	NwdafRegistration *NwdafRegistrationInfo        `json:"nwdafRegistration,omitempty"`
	Amf3Gpp           *Amf3GppAccessRegistration    `json:"amf3Gpp,omitempty"`
	AmfNon3Gpp        *AmfNon3GppAccessRegistration `json:"amfNon3Gpp,omitempty"`
	SmfRegistration   *SmfRegistrationInfo          `json:"smfRegistration,omitempty"`
	Smsf3Gpp          *SmsfRegistration             `json:"smsf3Gpp,omitempty"`
	SmsfNon3Gpp       *SmsfRegistration             `json:"smsfNon3Gpp,omitempty"`
	IpSmGw            *IpSmGwRegistration           `json:"ipSmGw,omitempty"`
}
