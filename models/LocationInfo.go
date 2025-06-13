/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:29 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type LocationInfo struct {
	Supi                         string                     `json:"supi,omitempty"`
	Gpsi                         string                     `json:"gpsi,omitempty"`
	RegistrationLocationInfoList []RegistrationLocationInfo `json:"registrationLocationInfoList"`
	SupportedFeatures            string                     `json:"supportedFeatures,omitempty"`
}
