/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 16:49:32 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProSeAuthenticationInfoResult struct {
	AuthType                   AuthType        `json:"authType"`
	ProseAuthenticationVectors []AvEapAkaPrime `json:"proseAuthenticationVectors,omitempty"`
	Supi                       string          `json:"supi,omitempty"`
	SupportedFeatures          string          `json:"supportedFeatures,omitempty"`
}
