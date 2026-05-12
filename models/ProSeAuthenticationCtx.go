/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Tue May 12 13:32:41 KST 2026 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProSeAuthenticationCtx struct {
	AuthType          AuthType        `json:"authType"`
	Links             map[string]Link `json:"_links"`
	ProSeAuthData     string          `json:"proSeAuthData"`
	SupportedFeatures string          `json:"supportedFeatures,omitempty"`
}
