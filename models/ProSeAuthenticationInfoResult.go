/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 11:41:44 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProSeAuthenticationInfoResult struct {
	SupportedFeatures          string          `json:"supportedFeatures,omitempty"`
	AuthType                   AuthType        `json:"authType"`
	ProseAuthenticationVectors []AvEapAkaPrime `json:"proseAuthenticationVectors,omitempty"`
	Supi                       string          `json:"supi,omitempty"`
}
