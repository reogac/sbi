/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jul 18 15:09:45 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProSeAuthenticationCtx struct {
	AuthType          AuthType        `json:"authType"`
	Links             map[string]Link `json:"_links"`
	ProSeAuthData     string          `json:"proSeAuthData"`
	SupportedFeatures string          `json:"supportedFeatures,omitempty"`
}
