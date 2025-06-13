/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Jun 13 13:39:25 KST 2025 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package models

type ProSeEapSession struct {
	EapPayload        string          `json:"eapPayload"`
	KnrProSe          string          `json:"knrProSe,omitempty"`
	Links             map[string]Link `json:"_links,omitempty"`
	AuthResult        AuthResult      `json:"authResult,omitempty"`
	SupportedFeatures string          `json:"supportedFeatures,omitempty"`
	Nonce2            string          `json:"nonce2,omitempty"`
	FiveGPrukId       string          `json:"5gPrukId,omitempty"`
}
