/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:51 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type Pp5gVnGroupProfileData struct {
	AllowedMtcProviders map[string]AllowedMtcProviderInfo `json:"allowedMtcProviders,omitempty"`
	SupportedFeatures   string                            `json:"supportedFeatures,omitempty"`
}
