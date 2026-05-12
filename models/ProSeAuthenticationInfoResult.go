/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:38 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProSeAuthenticationInfoResult struct {
	ProseAuthenticationVectors []AvEapAkaPrime `json:"proseAuthenticationVectors,omitempty"`
	Supi                       string          `json:"supi,omitempty"`
	SupportedFeatures          string          `json:"supportedFeatures,omitempty"`
	AuthType                   AuthType        `json:"authType"`
}
