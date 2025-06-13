/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:28:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type LocationInfo struct {
	RegistrationLocationInfoList []RegistrationLocationInfo `json:"registrationLocationInfoList"`
	SupportedFeatures            string                     `json:"supportedFeatures,omitempty"`
	Supi                         string                     `json:"supi,omitempty"`
	Gpsi                         string                     `json:"gpsi,omitempty"`
}
